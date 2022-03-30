package repository

import (
	"rafal-kalinowski.pl/domain/model"
	domain "rafal-kalinowski.pl/domain/service"
)

type purchaseRepository struct {
	repoConnector RepoConnector
}

func NewPurchaseRepository(c RepoConnector) domain.PurchaseRepository {
	repo := &purchaseRepository{
		repoConnector: c,
	}

	return repo
}

func (repo *purchaseRepository) GetPurchases(limit, offset uint) []model.PurchaseOrder {
	var purchase []model.PurchaseOrder

	count := 0
	model.DB.Find(&[]model.PurchaseOrder{}).Count(&count)

	model.DB.Preload("Product").
		Scopes(Paginate(int(limit), int(offset))).
		Find(&purchase)

	return purchase
}

func (repo *purchaseRepository) GetPurchase(id uint) (model.PurchaseOrder, error) {
	var purchase model.PurchaseOrder

	if err := model.DB.Where("id = ?", id).Preload("Product").
		First(&purchase, id).Error; err != nil {
		return purchase, err
	}

	return purchase, nil
}

func (repo *purchaseRepository) CreatePurchase(purchase *model.PurchaseOrder) error {
	if err := model.DB.Create(&purchase).Error; err != nil {
		return err
	}

	return nil
}

func (repo *purchaseRepository) GetPurchasesForUser(limit, offset uint, user string) []model.PurchaseOrder {
	var purchase []model.PurchaseOrder
	count := 0

	model.DB.Where("user_id = ?", user).Find(&[]model.PurchaseOrder{}).Count(&count)

	model.DB.Preload("Product").
		Scopes(Paginate(int(limit), int(offset))).
		Where("id = ?", user).Find(&[]model.PurchaseOrder{})

	model.DB.Where("user_id = ?", user).Preload("Product").Find(&purchase)
	return purchase
}
