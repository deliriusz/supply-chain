package domain

import (
	"bufio"

	"rafal-kalinowski.pl/domain/model"
)

type ProductService interface {
	GetProducts(limit, offset uint) ([]model.Product, uint)
	GetProduct(id uint) (model.Product, error)
	CreateProduct(*model.Product) error
	CreateImage(productId uint, file *bufio.Reader) (string, error)
	GetImage(fileName string) (string, string, *bufio.Reader, error)
}

type ProductRepository interface {
	GetProducts(limit, offset uint) ([]model.Product, uint)
	GetProduct(id uint) (model.Product, error)
	CreateProduct(*model.Product) error
	CreateImage(productId uint, file *bufio.Reader) (string, error)
	GetImage(fileName string) (string, string, *bufio.Reader, error)
}
