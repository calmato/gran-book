package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"runtime"

	"github.com/calmato/gran-book/infra/functions/gcr-cleaner/pkg/gcrcleaner"
	"github.com/google/go-containerregistry/pkg/v1/google"
)

var (
	serviceKey  = flag.String("token", os.Getenv("GCP_SERVICE_KEY_JSON"), "JSON which GCP service account")
	repo        = flag.String("repo", "", "Repository name")
	keep        = flag.Int("keep", 3, "Minimum to keep generation")
	allowTagged = flag.Bool("allow-tagged", false, "Delete tagged images")
	tagFilter   = flag.String("tag-filter", "", "Tags pattern to clean")
)

func main() error {
	flag.Parse()

	err := exec()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	return nil
}

func exec() error {
	if *serviceKey == "" {
		return fmt.Errorf("the environment variable GCP_SERVICE_KEY_JSON is required")
	}

	if *repo == "" {
		return fmt.Errorf("missing -repo")
	}

	if !*allowTagged && *tagFilter != "" {
		return fmt.Errorf("-allow-tagged must be true when -tag-filter is declared")
	}

	tagFilterRegexp, err := regexp.Compile(*tagFilter)
	if err != nil {
		return fmt.Errorf("failed to parse -tag-filter: %w", err)
	}

	auth := google.NewJSONKeyAuthenticator(*serviceKey)
	concurrency := runtime.NumCPU()
	cleaner, err := gcrcleaner.NewCleaner(auth, concurrency)
	if err != nil {
		return fmt.Errorf("failed to create cleaner: %w", err)
	}

	// Do the deletion.
	fmt.Fprintf(os.Stdout, "%s: deleting refs older than %d generations\n", *repo, *keep)

	deleted, err := cleaner.Clean(*repo, *keep, *allowTagged, tagFilterRegexp)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "%s: successfully deleted %d refs\n", *repo, len(deleted))

	return nil
}
