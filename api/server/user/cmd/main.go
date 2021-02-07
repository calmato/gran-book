package main

import (
	"github.com/calmato/gran-book/api/server/user/config"
)

func main() {
	err := config.Execute()
	if err != nil {
		panic(err)
	}
}
