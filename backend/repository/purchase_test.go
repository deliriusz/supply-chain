package repository_test

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/franela/goblin"
	"rafal-kalinowski.pl/domain/model"
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

		g.It("Should retrieve created product", func() {
			g.Timeout(time.Hour)
			purchase := createRandomPurchase()
			if err := purchaseRepo.CreatePurchase(purchase); err != nil {
				g.Fail(err)
			}

			dbPurchase, err := purchaseRepo.GetPurchase(1)
			if err != nil {
				g.Fail(err)
			}

			g.Assert(dbPurchase.Id == 1).IsTrue()
			g.Assert(dbPurchase.Date.Unix()).Equal(purchase.Date.Unix())
			g.Assert(dbPurchase.UserId).Equal(purchase.UserId)
			g.Assert(dbPurchase.Product[0].ModelId > 0).IsTrue()
			g.Assert(dbPurchase.Product[0].ModelId).Equal(purchase.Product[0].ModelId)

		})
	})
}

func createRandomPurchase() *model.PurchaseOrder {
	createRandomProductModel()
	createRandomProductModel()
	purchaseProduct1 := createRandomProduct(0)
	purchaseProduct2 := createRandomProduct(0)

	purchaseProduct2.Id += 2

	poTime := time.Now().Add(time.Hour * (time.Duration(-rand.Intn(5000))))
	return &model.PurchaseOrder{
		UserId: "asdf",
		Date:   poTime,
		Product: []model.Product{
			*purchaseProduct1,
			*purchaseProduct2,
		},
	}
}
