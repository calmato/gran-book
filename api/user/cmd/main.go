package main

import (
	"github.com/calmato/gran-book/api/user/config"
)

func main() {
	err := config.Execute()
	if err != nil {
		panic(err)
	}
}
