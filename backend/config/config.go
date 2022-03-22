package config

import (
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
)

var LOGIN_SESSION_TTL_IN_SECS int
var IMAGE_LOCAL_STORAGE string
var PROJECT_ROOT string
var ADDRESS_LOGIN_NONCE_MAP map[string]int64
var VALID_ADDRESS_REGEXP *regexp.Regexp

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	IMAGE_LOCAL_STORAGE = os.Getenv("IMAGE_LOCAL_STORAGE")
	PROJECT_ROOT = os.Getenv("PROJECT_ROOT")
	LOGIN_SESSION_TTL_IN_SECS, _ = strconv.Atoi(os.Getenv("LOGIN_SESSION_TTL_IN_SECS"))

	ADDRESS_LOGIN_NONCE_MAP = make(map[string]int64)
	VALID_ADDRESS_REGEXP = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
}
