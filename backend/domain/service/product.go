package domain

import (
	"bufio"

	"rafal-kalinowski.pl/domain/model"
)

type ProductService interface {
	GetProductModels(limit, offset uint) ([]model.ProductModel, uint)
	GetProductModel(id uint) (model.ProductModel, error)
	CreateProductModel(*model.ProductModel) error
	GetProduct(id uint) (*model.Product, error)
	GetProducts(limit, offset uint) (*[]model.Product, uint)
	GetProductsForUser(user string, limit, offset uint) (*[]model.Product, uint)
	CreateProduct(*model.Product) error
	ChangeProductState(id uint, state model.ProductState) error
	SellProduct(id uint, to string) error
	CreateImage(productId uint, file *bufio.Reader) (model.Image, error)
	GetImage(fileName string) (string, string, *bufio.Reader, error)
}

type ProductRepository interface {
	GetProductModels(limit, offset uint) ([]model.ProductModel, uint)
	GetProductModel(id uint) (model.ProductModel, error)
	CreateProductModel(*model.ProductModel) error
	GetProduct(id uint) (*model.Product, error)
	GetProducts(limit, offset uint) (*[]model.Product, uint)
	GetProductsForUser(user string, limit, offset uint) (*[]model.Product, uint)
	CreateProduct(*model.Product) error
	UpdateProduct(*model.Product) error
	CreateImage(productId uint, file *bufio.Reader) (model.Image, error)
	GetImage(fileName string) (string, string, *bufio.Reader, error)
}
