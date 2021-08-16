package main

import (
	"github.com/calmato/gran-book/api/gateway/admin/config"
)

func main() {
	config.CheckError(config.Execute())
}
