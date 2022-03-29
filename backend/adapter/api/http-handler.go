package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"rafal-kalinowski.pl/config"
	"rafal-kalinowski.pl/controller"
	"rafal-kalinowski.pl/domain/model"
	domain "rafal-kalinowski.pl/domain/service"
)

type HTTPHandler interface {
	domain.LoginService
	Start() error
	Init()
}

type httpHandler struct {
	loginService domain.LoginService
	router       *gin.Engine
}

func NewHTTPHandler(loginService domain.LoginService) *httpHandler {
	return &httpHandler{
		loginService: loginService,
	}
}

func (hdl *httpHandler) Start() error {
	return hdl.router.Run()
}

func (hdl *httpHandler) Init() {
	hdl.router = gin.Default()

	setupCors(hdl.router)
	setupAdminRoutes(hdl.router)
	setupAuthenticatedRoutes(hdl.router)
	hdl.setupNormalRoutes(hdl.router)
}

func (hdl *httpHandler) GetLoginChallenge(c *gin.Context) {
	var input model.LoginChallenge

	if err := checkLoginRequest(&input, c); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if login, err := hdl.loginService.GetLoginChallenge(&input); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"nonce": strconv.FormatInt(login.Nonce, 10)})
	}
}

func (hdl *httpHandler) Login(c *gin.Context) {
	var input model.LoginChallenge

	if err := checkLoginRequest(&input, c); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if login, err := hdl.loginService.Login(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Error(err)
		return
	} else {
		c.SetCookie(config.COOKIE_SESSIONID, login.SessionId, int(login.TTL), "/", "localhost", true, true)
	}
}

func (hdl *httpHandler) Logout(c *gin.Context) {
	cookie, _ := c.Cookie(config.COOKIE_SESSIONID)

	if err := hdl.loginService.Logout(&model.Login{
		SessionId: cookie,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie(config.COOKIE_SESSIONID, "", -1, "/", "localhost", true, true)
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

func (hdl *httpHandler) setupNormalRoutes(router *gin.Engine) {
	router.POST("/auth/challenge", hdl.GetLoginChallenge)
	router.POST("/auth/login", hdl.Login)
	router.GET("/auth/logout", hdl.Logout)
	router.GET("/product", controller.GetProducts)
	router.GET("/product/:id", controller.GetProduct)
	router.POST("/product/:id/image", controller.CreateImage)
	router.GET("/image/:fileName", controller.GetImage)
	router.POST("/purchase", controller.CreatePurchase)
}

func setupAuthenticatedRoutes(router *gin.Engine) {
	authenticatedRoutes := router.Group("/")
	{
		authenticatedRoutes.Use(authenticate(config.ROLE_CLIENT))
		authenticatedRoutes.GET("/purchase", controller.GetPurchases)
		authenticatedRoutes.GET("/purchase/:id", controller.GetPurchase)
		authenticatedRoutes.GET("/purchase/user/:id", controller.GetPurchaseForUser)
	}
}

func setupAdminRoutes(router *gin.Engine) {
	adminRoutes := router.Group("/")
	{
		adminRoutes.Use(authenticate(config.ROLE_ADMIN))
		adminRoutes.POST("/product", controller.CreateProduct)
	}
}

func setupCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     config.CORS_ALLOW_ORIGINS,
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
