package domain

import (
	"mime/multipart"
	"os"

	"rafal-kalinowski.pl/domain/model"
)

type ProductService interface {
	GetProducts(limit, offset uint) ([]model.Product, uint)
	GetProduct(id uint) (model.Product, error)
	CreateProduct(*model.Product) error
	CreateImage(productId uint, file *multipart.File) (string, error)
	GetImage(fileName string) (string, string, *os.File, error)
}

type ProductRepository interface {
	GetProducts(limit, offset uint) ([]model.Product, uint)
	GetProduct(id uint) (model.Product, error)
	CreateProduct(*model.Product) error
	CreateImage(productId uint, file *multipart.File) (string, error)
	GetImage(fileName string) (string, string, *os.File, error)
}
