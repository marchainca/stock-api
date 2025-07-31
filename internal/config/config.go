package config

import (
	"log"
	"os"
)

type Config struct {
	BaseURL  string
	User     string
	Password string
	Port     string
}

func Load() Config {
	cfg := Config{
		BaseURL:  getEnv("API_BASE_URL", "https://api.karenai.click/swechallenge"),
		User:     os.Getenv("API_USER"),
		Password: os.Getenv("API_PASSWORD"),
		Port:     getEnv("PORT", "8580"),
	}

	if cfg.User == "" || cfg.Password == "" {
		log.Fatal("API_USER and API_PASSWORD must be set")
	}
	return cfg
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
