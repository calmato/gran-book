package config

import (
	"github.com/rs/cors"
)

func setCors() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type", "Accept"},
	})

	return c
}
