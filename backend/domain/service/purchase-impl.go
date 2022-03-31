package domain

import "rafal-kalinowski.pl/domain/model"

type purchaseService struct {
	repository PurchaseRepository
}

func (s *purchaseService) GetPurchases(limit, offset uint) ([]model.PurchaseOrder, uint) {
	return s.repository.GetPurchases(limit, offset)
}

func (s *purchaseService) GetPurchase(id uint) (model.PurchaseOrder, error) {
	return s.repository.GetPurchase(id)
}

func (s *purchaseService) CreatePurchase(po *model.PurchaseOrder) error {
	return s.repository.CreatePurchase(po)
}

func (s *purchaseService) GetPurchasesForUser(limit, offset uint, user string) ([]model.PurchaseOrder, uint) {
	return s.repository.GetPurchasesForUser(limit, offset, user)
}

func NewPurchaseService(repository PurchaseRepository) PurchaseService {
	return &purchaseService{
		repository,
	}
}
