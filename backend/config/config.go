package config

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type AUTH_ROLE int

const (
	ROLE_ADMIN AUTH_ROLE = iota
	ROLE_USER
	ROLE_CLIENT
)

var LOGIN_SESSION_TTL_IN_SECS int
var IMAGE_LOCAL_STORAGE string
var PROJECT_ROOT string
var COOKIE_SESSIONID string
var CORS_ALLOW_ORIGINS []string
var ADDRESS_LOGIN_NONCE_MAP map[string]int64
var VALID_ADDRESS_REGEXP *regexp.Regexp

func Init(envFiles ...string) {
	err := godotenv.Load(envFiles...)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	IMAGE_LOCAL_STORAGE = os.Getenv("IMAGE_LOCAL_STORAGE")
	CORS_ALLOW_ORIGINS = strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"), ",")
	PROJECT_ROOT = os.Getenv("PROJECT_ROOT")
	LOGIN_SESSION_TTL_IN_SECS, _ = strconv.Atoi(os.Getenv("LOGIN_SESSION_TTL_IN_SECS"))

	ADDRESS_LOGIN_NONCE_MAP = make(map[string]int64)
	VALID_ADDRESS_REGEXP = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	COOKIE_SESSIONID = "SESSIONID"
}
