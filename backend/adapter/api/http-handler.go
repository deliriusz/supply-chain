package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rafal-kalinowski.pl/config"
	"rafal-kalinowski.pl/controller"
	"rafal-kalinowski.pl/domain/model"
	domain "rafal-kalinowski.pl/domain/service"
)

type HTTPHandler interface {
	domain.LoginService
	domain.ProductService
	domain.PurchaseService
	Start() error
	Init()
}

type httpHandler struct {
	loginService    domain.LoginService
	productService  domain.ProductService
	purchaseService domain.PurchaseService
	router          *gin.Engine
}

func NewHTTPHandler(loginService domain.LoginService,
	productService domain.ProductService,
	purchaseService domain.PurchaseService) *httpHandler {
	return &httpHandler{
		loginService:    loginService,
		productService:  productService,
		purchaseService: purchaseService,
	}
}

func (hdl *httpHandler) Start() error {
	return hdl.router.Run()
}

func (hdl *httpHandler) Init() {
	hdl.router = gin.Default()

	hdl.setupCors()
	hdl.setupAdminRoutes()
	hdl.setupAuthenticatedRoutes()
	hdl.setupNormalRoutes()
}

func checkLoginRequest(input *model.LoginChallenge, c *gin.Context) error {
	if err := c.ShouldBindJSON(input); err != nil {
		return err
	}

	address := input.Address

	if !config.VALID_ADDRESS_REGEXP.MatchString(address) {
		return fmt.Errorf("invalid address: %s", address)
	}

	return nil
}

func (hdl *httpHandler) authenticate(role config.AUTH_ROLE) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie(config.COOKIE_SESSIONID)
		session, err := hdl.loginService.GetSessionById(cookie)
		currentTimestamp := time.Now().UnixMilli()

		if err != nil || session.ExpiresAt == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}

		if currentTimestamp > session.ExpiresAt {
			hdl.loginService.Logout(session)
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

func (hdl *httpHandler) setupNormalRoutes() {
	router := hdl.router

	router.POST("/auth/challenge", hdl.GetLoginChallenge)
	router.POST("/auth/login", hdl.Login)
	router.GET("/auth/logout", hdl.Logout)
	router.GET("/product", controller.GetProducts)
	router.GET("/product/:id", controller.GetProduct)
	router.POST("/product/:id/image", controller.CreateImage)
	router.GET("/image/:fileName", controller.GetImage)
	router.POST("/purchase", controller.CreatePurchase)
}

func (hdl *httpHandler) setupAuthenticatedRoutes() {
	authenticatedRoutes := hdl.router.Group("/")
	{
		authenticatedRoutes.Use(hdl.authenticate(config.ROLE_CLIENT))
		authenticatedRoutes.GET("/purchase", controller.GetPurchases)
		authenticatedRoutes.GET("/purchase/:id", controller.GetPurchase)
		authenticatedRoutes.GET("/purchase/user/:id", controller.GetPurchaseForUser)
	}
}

func (hdl *httpHandler) setupAdminRoutes() {
	adminRoutes := hdl.router.Group("/")
	{
		adminRoutes.Use(hdl.authenticate(config.ROLE_ADMIN))
		adminRoutes.POST("/product", controller.CreateProduct)
	}
}

func (hdl *httpHandler) setupCors() {
	hdl.router.Use(cors.New(cors.Config{
		AllowOrigins:     config.CORS_ALLOW_ORIGINS,
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
