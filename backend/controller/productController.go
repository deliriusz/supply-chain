package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"rafal-kalinowski.pl/config"
	"rafal-kalinowski.pl/model"
)

func GetProducts(c *gin.Context) {
	var products []model.Product
	var productDtos []model.ProductDTO

	model.DB.Preload("Img").Preload("Specification").Find(&products)

	for _, returnedProduct := range products {
		productDtos = append(productDtos, model.ToProductDTO(returnedProduct))
	}

	c.JSON(http.StatusOK, productDtos)
}

func FindProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, productExists := getProduct(uint(productId))

	if !productExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, model.ToProductDTO(product))
}

func CreateProduct(c *gin.Context) {
	var input model.ProductDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := model.ToProduct(input)

	model.DB.Create(&product)

	for _, sd := range input.Specification {
		model.DB.Create(model.ToSpecification(sd))
	}

	c.String(http.StatusAccepted, "")
}

func CreateImage(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, productExists := getProduct(uint(productId))

	if !productExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	file, _, err := c.Request.FormFile("upload")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	image := model.Image{
		ProductId: uint(productId),
		//TODO: change
		ImageName: fmt.Sprintf("%d-%d.png", productId, time.Now().UnixNano()),
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)

	os.WriteFile(config.IMAGE_LOCAL_STORAGE+image.ImageName, buf.Bytes(), 0644)

	model.DB.Create(&image)
}

func GetImage(c *gin.Context) {
	fs := http.FileSystem(http.Dir(config.IMAGE_LOCAL_STORAGE))

	c.FileFromFS(c.Param("fileName"), fs)
}

func getProduct(productId uint) (model.Product, bool) {
	var product model.Product

	if err := model.DB.Where("id = ?",
		productId).Preload("Img").Preload("Specification").First(&product, productId).Error; err != nil {
		return model.Product{}, false
	}

	return product, true
}
