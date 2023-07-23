package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load config %s", err)
	}

	return Config{
		Database{
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
		},
		Server{
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		},
	}
}
