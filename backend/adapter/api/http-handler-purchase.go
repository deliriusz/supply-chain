package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"rafal-kalinowski.pl/domain/model"
)

func (hdl *httpHandler) GetPurchases(c *gin.Context) {
	var purchaseDtos []model.PurchaseOrderDTO

	limit, offset := safePaginationFromContext(c)
	purchase, count := hdl.purchaseService.GetPurchases(uint(limit), uint(offset))

	for _, returnedPurchase := range purchase {
		purchaseDtos = append(purchaseDtos, model.ToPurchaseDTO(returnedPurchase))
	}

	c.JSON(http.StatusOK, gin.H{"total": count, "purchase": purchaseDtos})
}

func (hdl *httpHandler) CreatePurchase(c *gin.Context) {
	//TODO: implement
	panic("to implement")
}

func (hdl *httpHandler) GetPurchase(c *gin.Context) {
	var purchaseDtos model.PurchaseOrderDTO

	id := safeStringToInt(c.Param("id"), -1)

	if id < 0 {
		log.Error("invalid purchaseId")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid purchaseId"})
		return
	}

	if purchase, err := hdl.purchaseService.GetPurchase(uint(id)); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		purchaseDtos = model.ToPurchaseDTO(purchase)
		c.JSON(http.StatusOK, gin.H{"total": 1, "purchase": purchaseDtos})
	}
}

func (hdl *httpHandler) GetPurchasesForUser(c *gin.Context) {
	var purchaseDtos []model.PurchaseOrderDTO

	userId := c.Param("id")
	limit, offset := safePaginationFromContext(c)
	purchase, count := hdl.purchaseService.GetPurchasesForUser(uint(limit), uint(offset), userId)

	for _, returnedPurchase := range purchase {
		purchaseDtos = append(purchaseDtos, model.ToPurchaseDTO(returnedPurchase))
	}

	c.JSON(http.StatusOK, gin.H{"total": count, "purchase": purchaseDtos})
}
