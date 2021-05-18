package main

import (
	"github.com/calmato/gran-book/api/server/information/config"
)

func main() {
	err := config.Execute()
	if err != nil {
		panic(err)
	}
}
