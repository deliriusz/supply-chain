package repository_test

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestCreatePurchase(t *testing.T) {
	g := Goblin(t)

	g.Describe("Test CreatePurchase", func() {
		g.JustBeforeEach(Cleanup)
		g.JustBeforeEach(Setup)
		g.After(Cleanup)

		g.It("Should create a purchase without an error", func() {
			if err := purchaseRepo.CreatePurchase(createRandomPurchase()); err != nil {
				g.Fail(err)
			}
		})
	})
}
