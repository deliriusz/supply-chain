package api_test

import (
	"testing"

	. "github.com/franela/goblin"
)

// model.DB.Create(&model.PurchaseOrder{
// 	Id:      1,
// 	UserId:  1,
// 	Product: model.Product{Id: 1},
// 	Price:   1590,
// 	Date:    "2020-02-02T01:20:15Z",
// 	Status:  "IN_PROGRESS",
// })

// model.DB.Create(&model.PurchaseOrder{
// 	Id:      2,
// 	UserId:  1,
// 	Product: model.Product{Id: 2},
// 	Price:   99,
// 	Date:    "2022-11-22T15:21:36Z",
// 	Status:  "SENT",
// })

func TestCreatePurchase(t *testing.T) {
	g := Goblin(t)
	g.Describe("LoginChallenge", func() {
		g.JustBeforeEach(Cleanup)
		g.JustBeforeEach(Setup)
		g.After(Cleanup)

	})
}
