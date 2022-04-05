package repository_test

import (
	"testing"

	. "github.com/franela/goblin"
	"rafal-kalinowski.pl/config"
	"rafal-kalinowski.pl/domain/model"
)

//TODO: think about implementing message signing for login testing
func TestLoginChallenge(t *testing.T) {
	g := Goblin(t)

	loginChallenge1 := model.LoginChallenge{Address: "0x482BC0fBA93cAdf4fC894D49730F8d19e2f359FD"}
	loginChallenge2 := model.LoginChallenge{Address: "0xcAdf4fC894D49730F8d19e2f35930F8dDFD30F8d"}

	g.Describe("Test GetLoginChallenge", func() {
		g.BeforeEach(Setup)
		g.AfterEach(Cleanup)

		g.It("Should return different nonce for each address", func() {
			resp1, err1 := loginRepo.GetLoginChallenge(&loginChallenge1)
			if err1 != nil {
				g.Fail(err1)
			}

			resp2, err2 := loginRepo.GetLoginChallenge(&loginChallenge2)
			if err2 != nil {
				g.Fail(err2)
			}

			g.Assert(resp1.Nonce > 0).IsTrue()
			g.Assert(resp2.Nonce > 0).IsTrue()
			g.Assert(resp1.Nonce == resp2.Nonce).IsFalse()

			g.Assert(config.ADDRESS_LOGIN_NONCE_MAP[loginChallenge1.Address]).Equal(resp1.Nonce)
			g.Assert(config.ADDRESS_LOGIN_NONCE_MAP[loginChallenge2.Address]).Equal(resp2.Nonce)
		})

		g.It("Should return different nonce for each invocation from the same address", func() {
			responses := make(map[int64]bool)
			nonceGenerationIterations := 1000

			for i := 0; i < nonceGenerationIterations; i++ {
				resp, err := loginRepo.GetLoginChallenge(&loginChallenge1)
				if err != nil {
					g.Fail(err)
				}

				g.Assert(resp.Nonce > 0).IsTrue()
				g.Assert(responses[resp.Nonce]).IsFalse()

				responses[resp.Nonce] = true
			}

			g.Assert(len(responses)).Equal(nonceGenerationIterations)
		})

		g.It("Should assign nonces only for accounts that requested it", func() {
			nonceGenerationIterations := 1000

			for i := 0; i < nonceGenerationIterations; i++ {
				resp1, err1 := loginRepo.GetLoginChallenge(&loginChallenge1)
				if err1 != nil {
					g.Fail(err1)
				}

				resp2, err2 := loginRepo.GetLoginChallenge(&loginChallenge2)
				if err2 != nil {
					g.Fail(err2)
				}

				g.Assert(resp1.Nonce > 0).IsTrue()
				g.Assert(resp2.Nonce > 0).IsTrue()
			}

			g.Assert(len(config.ADDRESS_LOGIN_NONCE_MAP)).Equal(2)
		})
	})
}
