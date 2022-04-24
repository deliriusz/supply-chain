package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rafal-kalinowski.pl/config"
	"rafal-kalinowski.pl/domain/model"
	domain "rafal-kalinowski.pl/domain/service"
)

type HTTPHandler interface {
	GetLoginChallenge(*gin.Context)
	Login(*gin.Context)
	Logout(*gin.Context)
	GetProducts(*gin.Context)
	GetProduct(*gin.Context)
	CreateProduct(*gin.Context)
	CreateImage(*gin.Context)
	GetImage(*gin.Context)
	GetPurchases(*gin.Context)
	GetPurchase(*gin.Context)
	CreatePurchase(*gin.Context)
	GetPurchasesForUser(*gin.Context)
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
	purchaseService domain.PurchaseService) HTTPHandler {
	return &httpHandler{
		loginService:    loginService,
		productService:  productService,
		purchaseService: purchaseService,
		router:          &gin.Engine{},
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

func (hdl *httpHandler) authenticate(role model.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie(config.COOKIE_SESSIONID)
		session, err := hdl.loginService.GetSessionById(cookie)
		currentTimestamp := time.Now().UnixMilli()

		if err != nil || session.ExpiresAt == 0 {
			abortAuthWithMessage(c, "Unauthorized")
		}

		if currentTimestamp > session.ExpiresAt {
			hdl.loginService.Logout(session)
			abortAuthWithMessage(c, "Token expired. Please log in again.")
		}

		assignedRole, err := hdl.loginService.GetUserRole(session.Address)
		if err != nil {
			abortAuthWithMessage(c, "An error occured. Please try again later.")
		}

		switch role {
		case model.Admin:
			if assignedRole.Role != model.Admin {
				abortAuthWithMessage(c, "You don't have required permissions to perform this action")
			}

		case model.DashboardViewer:
			if assignedRole.Role > model.DashboardViewer {
				abortAuthWithMessage(c, "You don't have required permissions to perform this action")
			}

		case model.Client:
			if assignedRole.Role > model.Client {
				abortAuthWithMessage(c, "Please log in in order to perform this action.")
			}

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
	router.GET("/product", hdl.GetProducts)
	router.GET("/product/:id", hdl.GetProduct)
	router.GET("/image/:fileName", hdl.GetImage)
	router.POST("/purchase", hdl.CreatePurchase)
}

func (hdl *httpHandler) setupAuthenticatedRoutes() {
	authenticatedRoutes := hdl.router.Group("/")
	{
		authenticatedRoutes.Use(hdl.authenticate(model.Client))
		authenticatedRoutes.GET("/auth/logout", hdl.Logout)
		authenticatedRoutes.GET("/purchase", hdl.GetPurchases)
		authenticatedRoutes.GET("/purchase/:id", hdl.GetPurchase)
		authenticatedRoutes.GET("/purchase/user/:id", hdl.GetPurchasesForUser)
	}
}

func (hdl *httpHandler) setupAdminRoutes() {
	adminRoutes := hdl.router.Group("/")
	{
		adminRoutes.Use(hdl.authenticate(model.Admin))
		adminRoutes.POST("/product", hdl.CreateProduct)
		adminRoutes.POST("/product/:id/image", hdl.CreateImage)
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

func safeStringToInt(numString string, defaultVal int) int {
	num, err := strconv.Atoi(numString)

	if err != nil {
		return defaultVal
	}

	return num
}

func safePaginationFromContext(c *gin.Context) (int, int) {
	safeLimit := safeStringToInt(c.Query("limit"), 10)
	safeOffset := safeStringToInt(c.Query("offset"), 0)

	if safeLimit < 0 {
		safeLimit = 10
	}

	if safeOffset < 0 {
		safeOffset = 0
	}

	return safeLimit, safeOffset
}

func abortAuthWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": message})
	c.Abort()
}
