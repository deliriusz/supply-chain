package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var IMAGE_LOCAL_STORAGE string
var PROJECT_ROOT string

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	IMAGE_LOCAL_STORAGE = os.Getenv("IMAGE_LOCAL_STORAGE")
	PROJECT_ROOT = os.Getenv("PROJECT_ROOT")
}
