package main

import (
	"github.com/calmato/gran-book/api/gateway/native/config"
)

func main() {
	err := config.Execute()
	if err != nil {
		panic(err)
	}
}
