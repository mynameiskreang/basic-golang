package connector

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	// Read these from environment variables or configuration files!
	OmisePublicKey = ""
	OmiseSecretKey = ""
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	OmisePublicKey = os.Getenv("OmisePublicKey")
	OmiseSecretKey = os.Getenv("OmiseSecretKey")
}
