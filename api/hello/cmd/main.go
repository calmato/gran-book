package main

import (
	"log"

	"github.com/calmato/gran-book/api/hello/config"
)

var (
	serverPort = "8080"
	logPath    = "/var/log/api"
)

func main() {
	gs, err := config.NewGRPCServer(serverPort, logPath)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	if err := gs.Serve(); err != nil {
		panic(err)
	}
}
