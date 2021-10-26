package main

import (
	"os"

	"github.com/calmato/gran-book/api/internal/book/cmd"
)

func main() {
	if err := cmd.Exec(); err != nil {
		os.Exit(1)
	}
}
