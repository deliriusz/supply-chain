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
var IMAGE_REPO_BASE_URI string
var PROJECT_ROOT string
var COOKIE_SESSIONID string
var CORS_ALLOW_ORIGINS []string
var ADDRESS_LOGIN_NONCE_MAP map[string]int64
var AUTHORIZATION_CONTRACT_ADDRESS string
var ETH_PROVIDER_URI string
var ETH_PRIVATE_KEY string
var VALID_ADDRESS_REGEXP *regexp.Regexp

func Init(envFiles ...string) {
	err := godotenv.Load(envFiles...)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	IMAGE_REPO_BASE_URI = os.Getenv("IMAGE_REPO_BASE_URI")
	IMAGE_LOCAL_STORAGE = os.Getenv("IMAGE_LOCAL_STORAGE")
	CORS_ALLOW_ORIGINS = strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"), ",")
	PROJECT_ROOT = os.Getenv("PROJECT_ROOT")
	LOGIN_SESSION_TTL_IN_SECS, _ = strconv.Atoi(os.Getenv("LOGIN_SESSION_TTL_IN_SECS"))
	AUTHORIZATION_CONTRACT_ADDRESS = os.Getenv("AUTHORIZATION_CONTRACT_ADDRESS")
	ETH_PROVIDER_URI = os.Getenv("ETH_PROVIDER_URI")
	ETH_PRIVATE_KEY = os.Getenv("ETH_PRIVATE_KEY")

	ADDRESS_LOGIN_NONCE_MAP = make(map[string]int64)
	VALID_ADDRESS_REGEXP = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	COOKIE_SESSIONID = "SESSIONID"
}
