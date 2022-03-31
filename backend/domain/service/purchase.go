package domain

import (
	"rafal-kalinowski.pl/domain/model"
)

type PurchaseService interface {
	GetPurchases(limit, offset uint) ([]model.PurchaseOrder, uint)
	GetPurchase(id uint) (model.PurchaseOrder, error)
	CreatePurchase(*model.PurchaseOrder) error
	GetPurchasesForUser(limit, offset uint, user string) ([]model.PurchaseOrder, uint)
}

type PurchaseRepository interface {
	GetPurchases(limit, offset uint) ([]model.PurchaseOrder, uint)
	GetPurchase(id uint) (model.PurchaseOrder, error)
	CreatePurchase(*model.PurchaseOrder) error
	GetPurchasesForUser(limit, offset uint, user string) ([]model.PurchaseOrder, uint)
}
