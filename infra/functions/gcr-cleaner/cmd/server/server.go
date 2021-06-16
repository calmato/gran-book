package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/calmato/gran-book/infra/functions/gcr-cleaner/pkg/gcrcleaner"
)

const (
	contentTypeHeader = "Content-Type"
	contentTypeJSON   = "application/json"
)

// Server - is a cleaning server.
type Server struct {
	cleaner *gcrcleaner.Cleaner
}

// NewServer - creates a new server for handler functions.
func NewServer(c *gcrcleaner.Cleaner) (*Server, error) {
	if c == nil {
		return nil, fmt.Errorf("missing cleaner")
	}

	s := &Server{
		cleaner: c,
	}

	return s, nil
}

// HTTPHandler - is an http handler that invokes the cleaner with the given parameters.
func (s *Server) HTTPHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deleted, status, err := s.clean(r.Body)
		if err != nil {
			s.handleError(w, err, status)
			return
		}

		b, err := json.Marshal(&cleanResp{
			Count: len(deleted),
			Refs:  deleted,
		})

		if err != nil {
			err = fmt.Errorf("failed to marshal JSON errors: %w", err)
			s.handleError(w, err, 500)
			return
		}

		w.WriteHeader(200)
		w.Header().Set(contentTypeHeader, contentTypeJSON)
		fmt.Fprint(w, string(b))
	}
}

// PubSubHandler is an http handler that invokes the cleaner from a pubsub
// request. Unlike an HTTP request, the pubsub endpoint always returns a success
// unless the pubsub message is malformed.
func (s *Server) PubSubHandler(cache Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := pubsubMessage{}

		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			err = fmt.Errorf("failed to decode pubsub message: %w", err)
			s.handleError(w, err, 400)
		}

		// PubSub is "at least once" delivery. The cleaner is idempotent,
		// but let's try to prevent unnecessary work by processing messages we've already received.
		msgID := fmt.Sprintf("%s/%s", m.Subscription, m.Message.ID)
		if exists := cache.Insert(msgID); exists {
			log.Printf("already processed message %s", msgID)
			w.WriteHeader(204)
			return
		}

		if len(m.Message.Data) == 0 {
			err := fmt.Errorf("missing data in pubsub payload")
			s.handleError(w, err, 400)
			return
		}

		// Start a goroutine to delete the images
		body := ioutil.NopCloser(bytes.NewReader(m.Message.Data))
		go func() {
			_, _, err := s.clean(body)
			if err != nil {
				log.Printf("error async: %s", err.Error())
			}
		}()

		w.WriteHeader(204)
	}
}

func (s *Server) handleError(w http.ResponseWriter, err error, status int) {
	log.Printf("error %d: %s", status, err.Error())

	b, err := json.Marshal(&errorResp{
		Error: err.Error(),
	})

	if err != nil {
		err = fmt.Errorf("failed to marshal JSON errors: %w", err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(status)
	w.Header().Set(contentTypeHeader, contentTypeJSON)
	fmt.Fprint(w, string(b))
}

func (s *Server) clean(r io.ReadCloser) ([]string, int, error) {
	p := Payload{}

	err := json.NewDecoder(r).Decode(&p)
	if err != nil {
		return nil, 500, fmt.Errorf("failed to decode payload as JSON: %w", err)
	}

	repo := p.Repo
	keep := p.Keep
	allowTagged := p.AllowTagged
	tagFilterRegexp, err := regexp.Compile(p.TagFilter)
	if err != nil {
		return nil, 500, fmt.Errorf("failed to parse tag_filter: %w", err)
	}

	// Do the deletion.
	log.Printf("%s: deleting refs older than %d generations\n", repo, keep)

	deleted, err := s.cleaner.Clean(repo, keep, allowTagged, tagFilterRegexp)
	if err != nil {
		return nil, 400, fmt.Errorf("failed to clean: %w", err)
	}

	log.Printf("delted %d refs for %s", len(deleted), repo)

	return deleted, 200, nil
}

// Payload - is the expected incoming payload format.
type Payload struct {
	// Repo is the name of the repo in the format gcr.io/foo/bar
	Repo string `json:"repo"`

	// Keep is the minimum number of images to keep.
	Keep int `json:"keep"`

	// AllowTagged is a Boolean value determine if tagged images are allowed
	// to be deleted.
	AllowTagged bool `json:"allow_tagged"`

	// TagFilter is the tags pattern to be allowed removing
	TagFilter string `json:"tag_filter"`
}

type cleanResp struct {
	Count int      `json:"count"`
	Refs  []string `json:"refs"`
}

type errorResp struct {
	Error string `json:"error"`
}

type pubsubMessage struct {
	Message struct {
		Data []byte `json:"data"`
		ID   string `json:"message_id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}
