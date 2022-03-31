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

func (r *purchaseRepository) GetPurchases(limit, offset uint) ([]model.PurchaseOrder, uint) {
	DB := r.repoConnector.GetConnector()
	var purchase []model.PurchaseOrder

	count := uint(0)
	DB.Find(&[]model.PurchaseOrder{}).Count(&count)

	DB.Preload("Product").
		Scopes(Paginate(int(limit), int(offset))).
		Find(&purchase)

	return purchase, count
}

func (r *purchaseRepository) GetPurchase(id uint) (model.PurchaseOrder, error) {
	DB := r.repoConnector.GetConnector()
	var purchase model.PurchaseOrder

	if err := DB.Where("id = ?", id).Preload("Product").
		First(&purchase, id).Error; err != nil {
		return purchase, err
	}

	return purchase, nil
}

func (r *purchaseRepository) CreatePurchase(purchase *model.PurchaseOrder) error {
	DB := r.repoConnector.GetConnector()
	if err := DB.Create(&purchase).Error; err != nil {
		return err
	}

	return nil
}

func (r *purchaseRepository) GetPurchasesForUser(limit, offset uint, user string) ([]model.PurchaseOrder, uint) {
	DB := r.repoConnector.GetConnector()
	var purchase []model.PurchaseOrder
	count := uint(0)

	DB.Where("user_id = ?", user).Find(&[]model.PurchaseOrder{}).Count(&count)

	DB.Preload("Product").
		Scopes(Paginate(int(limit), int(offset))).
		Where("id = ?", user).Find(&[]model.PurchaseOrder{})

	DB.Where("user_id = ?", user).Preload("Product").Find(&purchase)
	return purchase, count
}
