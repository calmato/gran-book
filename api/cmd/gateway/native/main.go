package main

import (
	"os"

	cmd "github.com/calmato/gran-book/api/internal/gateway/cmd/native"
)

func main() {
	if err := cmd.Exec(); err != nil {
		os.Exit(1)
	}
}
