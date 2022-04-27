package domain

import (
	"bufio"

	"rafal-kalinowski.pl/domain/model"
)

type ProductService interface {
	GetProductModels(limit, offset uint) ([]model.ProductModel, uint)
	GetProductModel(id uint) (model.ProductModel, error)
	CreateProductModel(*model.ProductModel) error
	CreateImage(productId uint, file *bufio.Reader) (model.Image, error)
	GetImage(fileName string) (string, string, *bufio.Reader, error)
}

type ProductRepository interface {
	GetProductModels(limit, offset uint) ([]model.ProductModel, uint)
	GetProductModel(id uint) (model.ProductModel, error)
	CreateProductModel(*model.ProductModel) error
	CreateImage(productId uint, file *bufio.Reader) (model.Image, error)
	GetImage(fileName string) (string, string, *bufio.Reader, error)
}
