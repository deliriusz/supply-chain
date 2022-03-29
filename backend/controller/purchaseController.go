package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"rafal-kalinowski.pl/domain/model"
)

func GetPurchases(c *gin.Context) {
	var purchase []model.PurchaseOrder
	var purchaseDtos []model.PurchaseOrderDTO

	count := 0
	model.DB.Find(&[]model.PurchaseOrder{}).Count(&count)

	model.DB.Preload("Product").
		Scopes(model.Paginate(c.Query("limit"), c.Query("offset"))).
		Find(&purchase)

	for _, returnedPurchase := range purchase {
		purchaseDtos = append(purchaseDtos, model.ToPurchaseDTO(returnedPurchase))
	}

	c.JSON(http.StatusOK, gin.H{"total": count, "purchase": purchaseDtos})
}

func CreatePurchase(c *gin.Context) {
	var purchase []model.PurchaseOrder
	var purchaseDtos []model.PurchaseOrderDTO

	count := 0
	model.DB.Find(&[]model.PurchaseOrder{}).Count(&count)

	model.DB.Preload("Product").
		Scopes(model.Paginate(c.Query("limit"), c.Query("offset"))).
		Find(&purchase)

	for _, returnedPurchase := range purchase {
		purchaseDtos = append(purchaseDtos, model.ToPurchaseDTO(returnedPurchase))
	}

	c.JSON(http.StatusOK, gin.H{"total": count, "purchase": purchaseDtos})
}

func GetPurchase(c *gin.Context) {
	var purchase model.PurchaseOrder
	var purchaseDtos model.PurchaseOrderDTO

	id := c.Param("id")

	if err := model.DB.Where("id = ?", id).Preload("Product").
		First(&purchase, id).Error; err != nil {
		log.Error(err)

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	purchaseDtos = model.ToPurchaseDTO(purchase)

	c.JSON(http.StatusOK, gin.H{"total": 1, "purchase": purchaseDtos})
}

func GetPurchaseForUser(c *gin.Context) {
	var purchase []model.PurchaseOrder
	var purchaseDtos []model.PurchaseOrderDTO

	userId := c.Param("id")
	count := 0

	model.DB.Where("user_id = ?", userId).Find(&[]model.PurchaseOrder{}).Count(&count)

	model.DB.Preload("Product").
		Scopes(model.Paginate(c.Query("limit"), c.Query("offset"))).
		Where("id = ?", userId).Find(&[]model.PurchaseOrder{})

	if err := model.DB.Where("user_id = ?", userId).Preload("Product").
		Find(&purchase).Error; err != nil {
		log.Error(err)

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	for _, returnedPurchase := range purchase {
		purchaseDtos = append(purchaseDtos, model.ToPurchaseDTO(returnedPurchase))
	}

	c.JSON(http.StatusOK, gin.H{"total": count, "purchase": purchaseDtos})
}
