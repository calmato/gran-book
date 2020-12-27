package main

import (
	"flag"

	"github.com/calmato/gran-book/infra/gateway/config"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := config.Execute(); err != nil {
		panic(err)
	}
}
