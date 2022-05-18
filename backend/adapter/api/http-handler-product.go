package api

import (
	"bufio"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"rafal-kalinowski.pl/domain/model"
)

type GetProductModelsResponse struct {
	Total    int                  `json:"total"`
	Products []model.ProductModel `json:"products"`
}

func (hdl *httpHandler) GetProductModels(c *gin.Context) {
	var productDtos []model.ProductModelDTO
	limit, offset := safePaginationFromContext(c)
	products, count := hdl.productService.GetProductModels(uint(limit), uint(offset))

	for _, returnedProduct := range products {
		productDtos = append(productDtos, model.ToProductModelDTO(returnedProduct))
	}

	c.JSON(http.StatusOK, gin.H{"total": count, "products": productDtos})
}

func (hdl *httpHandler) GetProductModel(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error(err)
		return
	}

	product, err := hdl.productService.GetProductModel(uint(productId))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, model.ToProductModelDTO(product))
}

func (hdl *httpHandler) CreateProductModel(c *gin.Context) {
	var input model.ProductModelDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error(err)
		return
	}

	product := model.ToProductModel(input)

	if err := hdl.productService.CreateProductModel(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": product.Id})
}

func (hdl *httpHandler) GetProductMetadata(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error(err)
		return
	}

	product, err := hdl.productService.GetProductMetadata(uint(productId))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, model.ToProductMetadataDTO(*product))
}

func (hdl *httpHandler) GetProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error(err)
		return
	}

	product, err := hdl.productService.GetProduct(uint(productId))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, model.ToProductDTO(*product))
}

func (hdl *httpHandler) GetProducts(c *gin.Context) {
	session, err := hdl.getSessionByIdFromCookie(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	if session.Role != model.Admin {
		abortAuthWithMessage(c, http.StatusForbidden, "you don't have permission to read info about other users' purchases")
		return
	}

	var productDtos []model.ProductDTO
	limit, offset := safePaginationFromContext(c)
	products, count := hdl.productService.GetProducts(uint(limit), uint(offset))

	for _, returnedProduct := range *products {
		productDtos = append(productDtos, model.ToProductDTO(returnedProduct))
	}

	c.JSON(http.StatusOK, gin.H{"total": count, "products": productDtos})
}

func (hdl *httpHandler) GetProductsForUser(c *gin.Context) {
	userId := c.Param("userId")

	var productDtos []model.ProductDTO
	limit, offset := safePaginationFromContext(c)
	products, count := hdl.productService.GetProductsForUser(userId, uint(limit), uint(offset))

	for _, returnedProduct := range *products {
		productDtos = append(productDtos, model.ToProductDTO(returnedProduct))
	}

	c.JSON(http.StatusOK, gin.H{"total": count, "products": productDtos})
}

func (hdl *httpHandler) CreateProduct(c *gin.Context) {
	var input model.ProductDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error(err)
		return
	}

	product := model.ToProduct(input)

	if err := hdl.productService.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": product.Id})
}

func (hdl *httpHandler) ChangeProductState(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error(err)
		return
	}

	stateId, err := strconv.Atoi(c.Param("stateId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error(err)
		return
	}

	hdl.productService.ChangeProductState(uint(productId), model.ProductState(stateId))
}

func (hdl *httpHandler) CreateImage(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, _, err := c.Request.FormFile("upload")

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileReader := bufio.NewReader(file)

	if image, err := hdl.productService.CreateImage(uint(productId), fileReader); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, model.ToImageDTO(image))
	}
}

func (hdl *httpHandler) GetImage(c *gin.Context) {
	if filename, dir, _, err := hdl.productService.GetImage(c.Param("fileName")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		fs := http.FileSystem(http.Dir(dir))
		c.FileFromFS(filename, fs)
	}
}
