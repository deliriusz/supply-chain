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

func (s *productService) GetProduct(id uint) (*model.Product, error) {
	return s.repository.GetProduct(id)
}

func (s *productService) CreateProduct(product *model.Product) error {
	return s.repository.CreateProduct(product)
}

func (s *productService) ChangeProductState(id uint, state model.ProductState) error {
	if product, err := s.repository.GetProduct(id); err != nil {
		return err
	} else {
		product.State = state
		return s.repository.UpdateProduct(product)
	}
}

func (s *productService) GetProductMetadata(id uint) (*model.ProductMetadata, error) {
	return s.repository.GetProductMetadata(id)
}

func (s *productService) SellProduct(id uint, to string) error {
	if product, err := s.repository.GetProduct(id); err != nil {
		return err
	} else {
		product.Owner = to
		return s.repository.UpdateProduct(product)
	}
}

func (s *productService) GetProducts(limit, offset uint) (*[]model.Product, uint) {
	return s.repository.GetProducts(limit, offset)
}
func (s *productService) GetProductsForUser(user string, limit, offset uint) (*[]model.Product, uint) {
	return s.repository.GetProductsForUser(user, limit, offset)
}

func NewProductService(repository ProductRepository) ProductService {
	return &productService{
		repository,
	}
}
