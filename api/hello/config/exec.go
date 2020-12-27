package config

// Execute - gRPC Serverの起動
func Execute() error {
	env, err := loadEnvironment()
	if err != nil {
		return err
	}

	// gRPC Serverの設定取得
	gs, err := newGRPCServer(env.Port, env.LogPath, env.LogLevel)
	if err != nil {
		return err
	}

	// メトリクス用のHTTP Serverの起動
	hs, err := newHTTPServer(env.MetricsPort)
	if err != nil {
		return err
	}

	go func() {
		if err := hs.Serve(); err != nil {
			panic(err)
		}
	}()

	// gRPC Serverの起動
	if err := gs.Serve(); err != nil {
		panic(err)
	}

	return nil
}
