package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file!!! Err: %s", err)
	}

	return os.Getenv("MONGOURI")

}
