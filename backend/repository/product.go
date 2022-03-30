package repository

import (
	"mime/multipart"
	"os"

	"rafal-kalinowski.pl/domain/model"
	domain "rafal-kalinowski.pl/domain/service"
)

type productRepository struct {
	repoConnector RepoConnector
}

func NewProductRepository(c RepoConnector) domain.ProductRepository {
	repo := &productRepository{
		repoConnector: c,
	}

	return repo
}

func (repo *productRepository) GetProducts(limit, offset uint) []model.Product {

}
func (repo *productRepository) GetProduct(id uint) (model.Product, error) {

}
func (repo *productRepository) CreateProduct(*model.Product) error {

}
func (repo *productRepository) CreateImage(productId uint, file *multipart.File) error {

}
func (repo *productRepository) GetImage(fileName string) (name, dir string, file *os.File) {

}
