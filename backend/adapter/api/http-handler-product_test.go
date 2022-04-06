package api_test

import (
	"testing"

	. "github.com/franela/goblin"
	"github.com/gin-gonic/gin"
	"rafal-kalinowski.pl/config"
)

// model.DB.Create(&model.Image{Id: 101, ProductId: 101, Data: []byte("")})
// model.DB.Create(&model.Image{Id: 102, ProductId: 101, Data: []byte("")})

// model.DB.Create(&model.Product{
// 	Id: 1,
// 	Img: []model.Image{{Id: 1, ProductId: 1, Data: []byte("")},
// 		{Id: 2, ProductId: 1, Data: []byte("")}},
// 	Title:             "title",
// 	Description:       "lorem ipslum",
// 	Price:             1000,
// 	AvailableQuantity: 20,
// 	Specification: []model.Specification{
// 		{Id: 1, ProductId: 1, Name: "size", Value: "6x2 inch"},
// 		{Id: 2, ProductId: 1, Name: "color", Value: "red"},
// 	},
// })

func TestCreateProduct(t *testing.T) {
	g := Goblin(t)
	config.Init("../.env")
	router := gin.Default()

	router.GET("/product", httpApi.GetProducts)
	router.GET("/product/:id", httpApi.GetProduct)
	router.POST("/product/:id/image", httpApi.CreateImage)
	router.GET("/image/:fileName", httpApi.GetImage)
	router.POST("/product", httpApi.CreateProduct)

	g.Describe("Product", func() {
		g.JustBeforeEach(Cleanup)
		g.JustBeforeEach(Setup)
		g.After(Cleanup)

		// router := gin.Default()
		// router.POST("/auth/challenge", controller.GetLoginChallenge)

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
