package config

import (
	"os"
	"os/signal"
	"syscall"
)

// Execute - gRPC Serverの起動
func Execute() error {
	env, err := loadEnvironment()
	if err != nil {
		return err
	}

	gs, err := newGRPCServer(env.Port, env.LogPath, env.LogLevel)
	if err != nil {
		return err
	}

	go func() {
		if err := gs.Serve(); err != nil {
			panic(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	sig := <-sigChan
	switch sig {
	case syscall.SIGINT, syscall.SIGTERM:
		gs.Stop()
	}

	return nil
}
