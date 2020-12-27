package main

import (
	"flag"

	"github.com/calmato/gran-gook/infra/gateway/config"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := config.Execute(); err != nil {
		panic(err)
	}
}
