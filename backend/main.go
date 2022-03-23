package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"rafal-kalinowski.pl/config"
	"rafal-kalinowski.pl/controller"
	"rafal-kalinowski.pl/model"
)

func before() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before")
		c.Next()
	}
}

func after() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		fmt.Println("after")
	}
}

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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     config.CORS_ALLOW_ORIGINS,
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authenticatedRoutes := router.Group("/")
	{
		authenticatedRoutes.Use(before())
		authenticatedRoutes.Use(after())
		authenticatedRoutes.GET("/purchase", controller.GetPurchases)
		authenticatedRoutes.GET("/purchase/:id", controller.GetPurchase)
		authenticatedRoutes.GET("/purchase/user/:id", controller.GetPurchaseForUser)
	}

	adminRoutes := router.Group("/")
	{
		adminRoutes.Use(before())
	}

	router.POST("/auth/challenge", controller.GetLoginChallenge)
	router.POST("/auth/login", controller.Login)
	router.GET("/product", controller.GetProducts)
	router.POST("/product", controller.CreateProduct)
	router.GET("/product/:id", controller.GetProduct)
	router.POST("/product/:id/image", controller.CreateImage)
	router.GET("/image/:fileName", controller.GetImage)
	// router.GET("/purchase", controller.GetPurchases)
	// router.GET("/purchase/:id", controller.GetPurchase)
	// router.GET("/purchase/user/:id", controller.GetPurchaseForUser)
	router.POST("/purchase", controller.CreatePurchase)

	router.Run()
}
