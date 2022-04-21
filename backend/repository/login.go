package repository

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"rafal-kalinowski.pl/config"
	"rafal-kalinowski.pl/domain/model"
	domain "rafal-kalinowski.pl/domain/service"
)

type loginRepository struct {
	repoConnector RepoConnector
}

// GetLoginChallenge implements domain.LoginRepository
func (r *loginRepository) GetLoginChallenge(login *model.LoginChallenge) (*model.LoginChallenge, error) {
	nonce, err := getSecureRandom()
	if err != nil {
		return nil, err
	}

	config.ADDRESS_LOGIN_NONCE_MAP[login.Address] = nonce

	resp := &model.LoginChallenge{
		Nonce: nonce,
	}

	return resp, nil
}

// Login implements domain.LoginRepository
func (r *loginRepository) Login(login *model.LoginChallenge) (*model.Login, error) {
	DB := r.repoConnector.GetConnector()
	expectedNonce := config.ADDRESS_LOGIN_NONCE_MAP[login.Address]

	if !verifySig(login.Address, login.Signature, []byte(strconv.FormatInt(expectedNonce, 10))) {
		return nil, fmt.Errorf("invalid signature or nonce value for an account")
	}

	currentTimestamp := time.Now().UnixMilli()
	sessionTTL := config.LOGIN_SESSION_TTL_IN_SECS

	sessionIdNonceSeed, _ := getSecureRandom()
	sessionIdSeed := fmt.Sprintf("%s-%d-%d", login.Signature, currentTimestamp, sessionIdNonceSeed)
	sessionId := hex.EncodeToString(crypto.Keccak256([]byte(sessionIdSeed)))

	loginSession := &model.Login{
		Address:   login.Address,
		SessionId: sessionId,
		ExpiresAt: currentTimestamp + int64(sessionTTL)*1000,
		TTL:       uint(sessionTTL),
	}

	// standard DB.Where did not work - value starting with "0X" is automatically treated as hex number
	// and sqlite DB complained about number too big, even after trying to wrap it into quotes
	if err := DB.Exec("DELETE FROM logins WHERE ADDRESS = \"?\"", login.Address).Error; err != nil {
		return nil, err
	}

	if err := DB.Create(&loginSession).Error; err != nil {
		return nil, err
	}

	return loginSession, nil
}

// Logout implements domain.LoginRepository
func (r *loginRepository) Logout(login *model.Login) error {
	DB := r.repoConnector.GetConnector()
	var session model.Login

	err := DB.Where("session_id = ?", login.SessionId).
		First(&session, login.SessionId).Error

	if err == nil && session.ExpiresAt > 0 {
		return DB.Delete(&model.Login{}, session.Id).Error
	}

	return nil
}

func (r *loginRepository) GetSessionById(sessionId string) (*model.Login, error) {
	DB := r.repoConnector.GetConnector()
	var login model.Login

	if err := DB.Where("session_id = ?",
		sessionId).First(&login, sessionId).Error; err != nil {
		return &login, err
	}

	return &login, nil
}

func NewLoginRepository(c RepoConnector) domain.LoginRepository {
	repo := &loginRepository{
		repoConnector: c,
	}

	return repo
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
