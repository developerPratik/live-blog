package config

import (
	"github.com/joho/godotenv"
	"os"
)

func Load() error {
	environment := os.Getenv("ENV")

	if environment == "" {
		environment = "local"
	}
	err := godotenv.Load(".env." + environment)

	if err != nil {
		return err
	}
	return nil
}
