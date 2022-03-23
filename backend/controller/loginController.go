package controller

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"rafal-kalinowski.pl/config"
	"rafal-kalinowski.pl/model"
)

func GetLoginChallenge(c *gin.Context) {
	var input model.LoginChallenge

	if err := checkLoginRequest(&input, c); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nonce, err := getSecureRandom()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	config.ADDRESS_LOGIN_NONCE_MAP[input.Address] = nonce

	c.JSON(http.StatusOK, gin.H{"nonce": strconv.FormatInt(nonce, 10)})
}

func Login(c *gin.Context) {
	var input model.LoginChallenge

	if err := checkLoginRequest(&input, c); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expectedNonce := config.ADDRESS_LOGIN_NONCE_MAP[input.Address]

	if !verifySig(input.Address, input.Signature, []byte(strconv.FormatInt(expectedNonce, 10))) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid signature or nonce value for an account"})
		log.Error("Invalid signature or nonce value. Expected nonce ", expectedNonce, ", received nonce ", input.Nonce)
		return
	}

	currentTimestamp := time.Now().UnixMilli()
	sessionTTL := config.LOGIN_SESSION_TTL_IN_SECS

	sessionIdNonceSeed, _ := getSecureRandom()
	sessionIdSeed := fmt.Sprintf("%s-%d-%d", input.Signature, currentTimestamp, sessionIdNonceSeed)
	sessionId := hex.EncodeToString(crypto.Keccak256([]byte(sessionIdSeed)))

	loginSession := model.Login{
		Address:   input.Address,
		SessionId: sessionId,
		ExpiresAt: currentTimestamp + int64(sessionTTL)*1000,
	}

	err := model.DB.Create(&loginSession).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Error(err)
		return
	}

	c.SetCookie(config.COOKIE_SESSIONID, sessionId, sessionTTL, "/", "localhost", true, true)
}

func Logout(c *gin.Context) {
	cookie, _ := c.Cookie(config.COOKIE_SESSIONID)
	session, err := GetSessionById(cookie)

	if err == nil && session.ExpiresAt > 0 {
		RemoveSession(session.Id)
		c.SetCookie(config.COOKIE_SESSIONID, "", -1, "/", "localhost", true, true)
	}
}

func GetSessionById(sessionId string) (model.Login, error) {
	var login model.Login

	err := model.DB.Where("session_id = ?", sessionId).
		First(&login, sessionId).Error

	if err != nil {
		log.Error(err)
	}

	return login, err
}

func RemoveSession(id uint) error {
	err := model.DB.Delete(&model.Login{}, id).Error

	return err
}

func checkLoginRequest(input *model.LoginChallenge, c *gin.Context) error {
	if err := c.ShouldBindJSON(input); err != nil {
		return err
	}

	address := input.Address

	if !config.VALID_ADDRESS_REGEXP.MatchString(address) {
		return fmt.Errorf("invalid address: %s", address)
	}

	return nil
}

//https://gist.github.com/dcb9/385631846097e1f59e3cba3b1d42f3ed#file-eth_sign_verify-go
func verifySig(from, sigHex string, msg []byte) bool {
	fromAddr := common.HexToAddress(from)

	sig := hexutil.MustDecode(sigHex)
	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	if len(sig) != crypto.SignatureLength {
		return false
	}
	if sig[crypto.RecoveryIDOffset] != 27 && sig[crypto.RecoveryIDOffset] != 28 {
		return false
	}
	sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	pubKey, err := crypto.SigToPub(signHash(msg), sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return fromAddr == recoveredAddr
}

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func getSecureRandom() (int64, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64-1))
	if err != nil {
		return 1, err
	}
	return nBig.Int64() + 1, nil
}
