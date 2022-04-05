package domain

import (
	"bufio"

	"rafal-kalinowski.pl/domain/model"
)

type productService struct {
	repository ProductRepository
}

func (s *productService) GetProducts(limit, offset uint) ([]model.Product, uint) {
	return s.repository.GetProducts(limit, offset)
}

func (s *productService) GetProduct(id uint) (model.Product, error) {
	return s.repository.GetProduct(id)
}

func (s *productService) CreateProduct(product *model.Product) error {
	return s.repository.CreateProduct(product)
}

func (s *productService) CreateImage(productId uint, file *bufio.Reader) (string, error) {
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
