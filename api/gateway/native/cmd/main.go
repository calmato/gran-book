package main

import (
	"github.com/calmato/gran-book/api/gateway/native/config"
)

func main() {
	config.CheckError(config.Execute())
}
