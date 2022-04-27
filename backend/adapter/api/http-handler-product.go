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
		productDtos = append(productDtos, model.ToProductDTO(returnedProduct))
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

	c.JSON(http.StatusOK, model.ToProductDTO(product))
}

func (hdl *httpHandler) CreateProduct(c *gin.Context) {
	var input model.ProductModelDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error(err)
		return
	}

	product := model.ToProduct(input)

	if err := hdl.productService.CreateProductModel(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": product.Id})
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
