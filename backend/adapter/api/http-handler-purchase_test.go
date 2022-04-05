package api_test

import (
	"testing"

	. "github.com/franela/goblin"
	"github.com/gin-gonic/gin"
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
		router := gin.Default()
		router.POST("/auth/challenge", httpApi.GetLoginChallenge)

		// req, _ := http.NewRequest("POST", "/auth/challenge", nil)
		// Passing Test
		g.It("Should add two numbers ", func() {
			g.Assert(1 + 1).Equal(2)
		})
		// Failing Test
		g.It("Should match equal numbers", func() {
			g.Assert(2).Equal(4)
		})
		// Pending Test
		g.It("Should substract two numbers")
		// Excluded Test
		g.Xit("Should add two numbers ", func() {
			g.Assert(3 + 1).Equal(4)
		})
	})
}
