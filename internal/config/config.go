package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URL string
	PORT         string
}

func Load() *Config {
	godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	log.Printf("URL lida do ambiente: [%s]", dbURL)

	return &Config{
		DATABASE_URL: dbURL,
		PORT:         port,
	}
}
