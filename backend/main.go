package main

import (
	"github.com/gin-gonic/gin"

	"rafal-kalinowski.pl/config"
	"rafal-kalinowski.pl/controller"
	"rafal-kalinowski.pl/model"
)

func main() {
	router := gin.Default()
	config.Init()
	model.ConnectDatabase()

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

	router.GET("/product", controller.GetProducts)
	router.POST("/product", controller.CreateProduct)
	router.GET("/product/:id", controller.FindProduct)
	router.POST("/product/:id/image", controller.CreateImage)
	router.GET("/image/:fileName", controller.GetImage)

	router.Run()
}
