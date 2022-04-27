package domain

import (
	"bufio"

	"rafal-kalinowski.pl/domain/model"
)

type ProductService interface {
	GetProducts(limit, offset uint) ([]model.ProductModel, uint)
	GetProduct(id uint) (model.ProductModel, error)
	CreateProduct(*model.ProductModel) error
	CreateImage(productId uint, file *bufio.Reader) (model.Image, error)
	GetImage(fileName string) (string, string, *bufio.Reader, error)
}

type ProductRepository interface {
	GetProducts(limit, offset uint) ([]model.ProductModel, uint)
	GetProduct(id uint) (model.ProductModel, error)
	CreateProduct(*model.ProductModel) error
	CreateImage(productId uint, file *bufio.Reader) (model.Image, error)
	GetImage(fileName string) (string, string, *bufio.Reader, error)
}
