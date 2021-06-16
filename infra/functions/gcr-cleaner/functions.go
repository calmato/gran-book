package main

import (
	"fmt"
	"os"

	"github.com/calmato/gran-book/infra/functions/gcr-cleaner/cmd/cli"
)

func main() error {
	err := cli.Exec()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	return nil
}
