package api_test

import (
	"encoding/json"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gin-gonic/gin"
	"rafal-kalinowski.pl/config"
	"rafal-kalinowski.pl/adapter/api"
	"rafal-kalinowski.pl/domain/model"
)

func TestLoginChallenge(t *testing.T) {
	g := Goblin(t)
	config.Init("../../.env")
	router := gin.Default()
	LOGIN_CHALLENGE_URI := "/auth/challenge"

	router.POST(LOGIN_CHALLENGE_URI, api.GetLoginChallenge)

	validLoginChallenge := model.LoginChallenge{Address: "0x482BC0fBA93cAdf4fC894D49730F8d19e2f359FD"}
	stringifiedValidLoginChallenge, _ := json.Marshal(validLoginChallenge)

	invalidLoginChallenge := model.LoginChallenge{Address: "0xC0fBA93cAdf4fC894973482B"}
	stringifiedInvalidLoginChallenge, _ := json.Marshal(invalidLoginChallenge)

	g.Describe("Test LoginChallenge", func() {
		g.It("Should fail on invalid address", func() {
			respRecorder := api.ServeTestRequest(router, "POST", LOGIN_CHALLENGE_URI, stringifiedInvalidLoginChallenge)

			g.Assert(respRecorder.Code).Equal(http.StatusBadRequest)
			g.Assert(len(config.ADDRESS_LOGIN_NONCE_MAP)).Equal(0)
		})

		g.It("Should fail on wrong request", func() {
			respRecorder := api.ServeTestRequest(router, "POST", LOGIN_CHALLENGE_URI, nil)

			g.Assert(http.StatusBadRequest).Equal(respRecorder.Code)
			g.Assert(0).Equal(len(config.ADDRESS_LOGIN_NONCE_MAP))
		})

		g.It("Should generate random nonce on each valid request", func() {
			respRecorder := api.ServeTestRequest(router, "POST", LOGIN_CHALLENGE_URI, stringifiedValidLoginChallenge)
			g.Assert(http.StatusOK).Equal(respRecorder.Code)
			resp := respRecorder.Body.Bytes()
			loginChallengeResponse := model.LoginChallenge{}

			if err := json.Unmarshal(resp, &loginChallengeResponse); err != nil {
				g.Fail(err)
			}

			nonce := loginChallengeResponse.Nonce

			g.Assert(nonce).IsNotZero()
			g.Assert(nonce).Equal(config.ADDRESS_LOGIN_NONCE_MAP[validLoginChallenge.Address])

			// second nonce retrieval
			respRecorder = api.ServeTestRequest(router, "POST", LOGIN_CHALLENGE_URI, stringifiedValidLoginChallenge)
			g.Assert(http.StatusOK).Equal(respRecorder.Code)
			resp = respRecorder.Body.Bytes()

			if err := json.Unmarshal(resp, &loginChallengeResponse); err != nil {
				g.Fail(err)
			}

			secondNonce := loginChallengeResponse.Nonce
			g.Assert(secondNonce).IsNotZero()
			g.Assert(secondNonce).Equal(config.ADDRESS_LOGIN_NONCE_MAP[validLoginChallenge.Address])
			g.Assert(secondNonce != nonce).IsTrue("First and second nonce shold not be equal")
		})
	})
}
