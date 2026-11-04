package config

import "os"

type Config struct {
	DATABASE_URL string
	PORT         string
}

func Load() *Config {
	dbURL := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	return &Config{
		DATABASE_URL: dbURL,
		PORT:         port,
	}
}
