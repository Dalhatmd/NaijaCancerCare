package config

import (
	"os"
)

type Config struct {
	AppPort		string 
	DBUrl 		string
	JWTKey		string
}

func Load() *Config {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":8080"
	}
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = "default-url"
	}
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		key = "default-key"
	}
	return &Config{
		AppPort: port,
		DBUrl: url,
		JWTKey: key,
	}
}
