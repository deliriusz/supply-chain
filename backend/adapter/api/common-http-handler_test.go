package api_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	"rafal-kalinowski.pl/adapter/api"
	"rafal-kalinowski.pl/config"
	domain "rafal-kalinowski.pl/domain/service"
	"rafal-kalinowski.pl/repository"
)

var TABLE_NAME string
var loginRepository domain.LoginRepository
var productRepository domain.ProductRepository
var purchaseRepository domain.PurchaseRepository
var loginService domain.LoginService
var productService domain.ProductService
var purchaseService domain.PurchaseService
var httpApi api.HTTPHandler
var router *gin.Engine

func init() {
	TABLE_NAME = "firmex-api-test.db"
	config.Init("../../.env")
}

func Setup() {
	repoConnector := repository.NewRepoConnector()
	if err := repoConnector.InitConnection(TABLE_NAME, ""); err != nil {
		panic(err)
	}

	loginRepository = repository.NewLoginRepository(repoConnector)
	loginService = domain.NewLoginService(loginRepository)
	productRepository = repository.NewProductRepository(repoConnector)
	productService = domain.NewProductService(productRepository)
	purchaseRepository = repository.NewPurchaseRepository(repoConnector)
	purchaseService = domain.NewPurchaseService(purchaseRepository)
	httpApi = api.NewHTTPHandler(loginService, productService, purchaseService)

	router = gin.Default()
	//login
	router.POST("/auth/challenge", httpApi.GetLoginChallenge)

	//product
	router.GET("/product", httpApi.GetProducts)
	router.GET("/product/:id", httpApi.GetProduct)
	router.POST("/product/:id/image", httpApi.CreateImage)
	router.GET("/image/:fileName", httpApi.GetImage)
	router.POST("/product", httpApi.CreateProduct)

	//purchase
}

func Cleanup() {
	os.Remove("./" + TABLE_NAME)
	dir, _ := ioutil.ReadDir(config.IMAGE_LOCAL_STORAGE)
	for _, d := range dir {
		os.RemoveAll(config.IMAGE_LOCAL_STORAGE + d.Name())
	}
}

func ServeTestRequest(router *gin.Engine, method, uri string, data []byte, headers map[string]string) *httptest.ResponseRecorder {
	respRecorder := httptest.NewRecorder()
	req, _ := http.NewRequest(method, uri, bytes.NewReader(data))

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	router.ServeHTTP(respRecorder, req)

	return respRecorder
}
