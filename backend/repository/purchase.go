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

func (repo *purchaseRepository) GetPurchases(limit, offset uint) []model.PurchaseOrder {}
func (repo *purchaseRepository) GetPurchase(id uint) (model.PurchaseOrder, error)      {}
func (repo *purchaseRepository) CreatePurchase(*model.PurchaseOrder) error {

}
func (repo *purchaseRepository) GetPurchasesForUser(limit, offset uint, user string) []model.PurchaseOrder {
}
