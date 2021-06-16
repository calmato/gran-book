package gcrcleaner

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/gammazero/workerpool"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/google"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

type Cleaner struct {
	auth        authn.Authenticator
	concurrency int
}

// NewCleaner - Create a new GCR cleaner with the given token provider and concurrency.
func NewCleaner(auth authn.Authenticator, concurrency int) (*Cleaner, error) {
	c := &Cleaner{
		auth:        auth,
		concurrency: concurrency,
	}

	return c, nil
}

// Clean - Deletes old images from GCR.
func (c *Cleaner) Clean(repo string, keep int, allowTagged bool, tagFilterRegexp *regexp.Regexp) ([]string, error) {
	gcrrepo, err := name.NewRepository(repo)
	if err != nil {
		return nil, fmt.Errorf("failed to get repo %s: %w", repo, err)
	}

	tags, err := google.List(gcrrepo, google.WithAuth(c.auth))
	if err != nil {
		return nil, fmt.Errorf("failed to list tags for repo %s: %w", repo, err)
	}

	// Create a worker pool for parallel deletion
	pool := workerpool.New(c.concurrency)

	keepCount := 0
	deleted := make([]string, 0, len(tags.Manifests))
	deletedLock := sync.Mutex{}
	errs := map[string]error{}
	errsLock := sync.RWMutex{}

	manifests := make([]manifest, 0, len(tags.Manifests))
	for k, m := range tags.Manifests {
		manifests = append(manifests, manifest{k, m})
	}

	// Sort manifest by Created from the most recent to the least
	sort.Slice(manifests, func(i, j int) bool {
		return manifests[j].Info.Created.Before(manifests[i].Info.Created)
	})

	for _, m := range manifests {
		if c.shouldDelete(m.Info, allowTagged, tagFilterRegexp) {
			// Keep a certain amount of images
			if keepCount < keep {
				keepCount++
				continue
			}

			// Deletes all tags before deleting the image
			for _, tag := range m.Info.Tags {
				tagged := gcrrepo.Tag(tag)

				err := c.deleteOne(tagged)
				if err != nil {
					return nil, fmt.Errorf("failed to delete %s: %w", tagged, err)
				}
			}

			ref := gcrrepo.Digest(m.Digest)

			pool.Submit(func() {
				// Do not process if previous invocations failed.
				// This prevents a large build-up of failed requests and rate limit exceeding (e.g. bad auth).
				errsLock.RLock()
				if len(errs) > 0 {
					errsLock.RUnlock()
					return
				}
				errsLock.RUnlock()

				err := c.deleteOne(ref)
				if err != nil {
					cause := errors.Unwrap(err).Error()

					errsLock.Lock()
					if _, ok := errs[cause]; !ok {
						errs[cause] = err
						errsLock.RUnlock()
						return
					}
					errsLock.RUnlock()
				}

				deletedLock.Lock()
				deleted = append(deleted, m.Digest)
				deletedLock.Unlock()
			})
		}
	}

	// Wait for everything to finish
	pool.StopWait()

	// Aggregate any errors
	if len(errs) > 0 {
		errStrs := []string{}
		for _, err := range errs {
			errStrs = append(errStrs, err.Error())
		}

		if len(errStrs) == 1 {
			return nil, fmt.Errorf(errStrs[0])
		}

		return nil, fmt.Errorf("%d errors occurred: %s", len(errStrs), strings.Join(errStrs, ", "))
	}

	return deleted, nil
}

type manifest struct {
	Digest string
	Info   google.ManifestInfo
}

// deleteOne - Deletes a single repo ref using the supplied auth.
func (c *Cleaner) deleteOne(ref name.Reference) error {
	err := remote.Delete(ref, remote.WithAuth(c.auth))
	if err != nil {
		return fmt.Errorf("failed to delete %s: %w", ref, err)
	}

	return nil
}

// shouldDelete - Returns true if the manifest has no tags or allows deletion of tagged images.
func (c *Cleaner) shouldDelete(m google.ManifestInfo, allowTagged bool, tagFilterRegexp *regexp.Regexp) bool {
	return (len(m.Tags) == 0 || (allowTagged && tagFilterRegexp.MatchString(m.Tags[0])))
}
