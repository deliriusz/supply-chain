package domain

import (
	"bufio"

	"rafal-kalinowski.pl/domain/model"
)

type productService struct {
	repository ProductRepository
}

func (s *productService) GetProductModels(limit, offset uint) ([]model.ProductModel, uint) {
	return s.repository.GetProductModels(limit, offset)
}

func (s *productService) GetProductModel(id uint) (model.ProductModel, error) {
	return s.repository.GetProductModel(id)
}

func (s *productService) CreateProductModel(product *model.ProductModel) error {
	return s.repository.CreateProductModel(product)
}

func (s *productService) CreateImage(productId uint, file *bufio.Reader) (model.Image, error) {
	return s.repository.CreateImage(productId, file)
}

func (s *productService) GetImage(fileName string) (string, string, *bufio.Reader, error) {
	return s.repository.GetImage(fileName)
}

func NewProductService(repository ProductRepository) ProductService {
	return &productService{
		repository,
	}
}
