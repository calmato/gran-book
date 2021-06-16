package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/calmato/gran-book/infra/functions/gcr-cleaner/pkg/gcrcleaner"
	"github.com/google/go-containerregistry/pkg/v1/google"
)

func main() {
	// Disable timestamps in go logs because stackdriver has them already.
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serviceKey := os.Getenv("GCP_SERVICE_KEY_JSON")
	if serviceKey == "" {
		log.Fatalf("the environment variable GCP_SERVICE_KEY_JSON is requried")
	}

	auth := google.NewJSONKeyAuthenticator(serviceKey)
	concurrency := runtime.NumCPU()
	cleaner, err := gcrcleaner.NewCleaner(auth, concurrency)
	if err != nil {
		log.Fatalf("failed to create cleaner: %s", err)
	}

	cleanerServer, err := NewServer(cleaner)
	if err != nil {
		log.Fatalf("failed to create server: %s", err)
	}

	cache := NewTimerCache(5 * time.Minute)

	mux := http.NewServeMux()
	mux.Handle("/http", cleanerServer.HTTPHandler())
	mux.Handle("/pubsub", cleanerServer.PubSubHandler(cache))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	go func() {
		log.Printf("server is listening on %s\n", port)

		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("server exited: %s", err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	<-signalCh

	log.Printf("received stop, shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown server: %s", err)
	}
}
