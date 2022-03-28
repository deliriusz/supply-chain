package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"rafal-kalinowski.pl/config"
	"rafal-kalinowski.pl/controller"
	"rafal-kalinowski.pl/model"
)

func authenticate(role config.AUTH_ROLE) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie(config.COOKIE_SESSIONID)
		session, err := controller.GetSessionById(cookie)
		currentTimestamp := time.Now().UnixMilli()

		if err != nil || session.ExpiresAt == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}

		if currentTimestamp > session.ExpiresAt {
			controller.RemoveSession(session.Id)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired. Please log in again."})
			c.Abort()
		}

		switch role {
		case config.ROLE_ADMIN:
			//TODO: use smart contract to check if it's admin

		case config.ROLE_USER:
			//TODO: use smart contract to check if it's user

		default:
			//good to go, nothing to check now
		}

		c.Next()
	}
}

func main() {
	config.Init()
	model.ConnectDatabase()
	router := GetRouter()

	router.Run()
}

func GetRouter() *gin.Engine {
	router := gin.Default()

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
		authenticatedRoutes.Use(authenticate(config.ROLE_CLIENT))
		authenticatedRoutes.GET("/purchase", controller.GetPurchases)
		authenticatedRoutes.GET("/purchase/:id", controller.GetPurchase)
		authenticatedRoutes.GET("/purchase/user/:id", controller.GetPurchaseForUser)
	}

	adminRoutes := router.Group("/")
	{
		adminRoutes.Use(authenticate(config.ROLE_ADMIN))
	}

	router.POST("/auth/challenge", controller.GetLoginChallenge)
	router.POST("/auth/login", controller.Login)
	router.GET("/auth/logout", controller.Logout)
	router.GET("/product", controller.GetProducts)
	router.POST("/product", controller.CreateProduct)
	router.GET("/product/:id", controller.GetProduct)
	router.POST("/product/:id/image", controller.CreateImage)
	router.GET("/image/:fileName", controller.GetImage)
	// router.GET("/purchase", controller.GetPurchases)
	// router.GET("/purchase/:id", controller.GetPurchase)
	// router.GET("/purchase/user/:id", controller.GetPurchaseForUser)
	router.POST("/purchase", controller.CreatePurchase)

	return router
}
