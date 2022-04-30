package repository

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"rafal-kalinowski.pl/config"
	"rafal-kalinowski.pl/domain/model"
	domain "rafal-kalinowski.pl/domain/service"
)

type productRepository struct {
	repoConnector RepoConnector[*DBRepoConnector]
}

func NewProductRepository(c RepoConnector[*DBRepoConnector]) domain.ProductRepository {
	repo := &productRepository{
		repoConnector: c,
	}

	return repo
}

func (r *productRepository) GetProductModels(limit, offset uint) ([]model.ProductModel, uint) {
	DB := r.repoConnector.GetConnector().DB
	var products []model.ProductModel

	count := uint(0)
	DB.Find(&[]model.ProductModel{}).Count(&count)

	DB.Preload("Img").Preload("Specification").
		Scopes(Paginate(int(limit), int(offset))).
		Find(&products)

	return products, count
}

func (r *productRepository) GetProductModel(id uint) (model.ProductModel, error) {
	product, err := r.getProductModel(id)

	return product, err
}

func (r *productRepository) CreateProductModel(product *model.ProductModel) error {
	DB := r.repoConnector.GetConnector().DB
	if err := DB.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) GetProduct(id uint) (*model.Product, error) {
	DB := r.repoConnector.GetConnector().DB

	var product model.Product

	if err := DB.Where("id = ?",
		id).Preload("Model").First(&product, id).Error; err != nil {
		return &model.Product{}, err
	}

	return &product, nil
}

func (r *productRepository) GetProducts(limit, offset uint) (*[]model.Product, uint) {
	DB := r.repoConnector.GetConnector().DB
	var products []model.Product

	count := uint(0)
	DB.Find(&[]model.Product{}).Count(&count)

	DB.Preload("Model").
		Scopes(Paginate(int(limit), int(offset))).
		Find(&products)

	return &products, count
}

func (r *productRepository) GetProductsForUser(user string, limit, offset uint) (*[]model.Product, uint) {
	DB := r.repoConnector.GetConnector().DB
	var product *[]model.Product
	count := uint(0)

	DB.Where("owner = ?", user).Find(&[]model.Product{}).Count(&count)

	DB.Preload("Model").
		Scopes(Paginate(int(limit), int(offset))).
		Where("user_id = ?", user).Find(product)

	// DB.Where("user_id = ?", user).Preload("Product").Find(product)
	return product, count
}

func (r *productRepository) CreateProduct(product *model.Product) error {
	DB := r.repoConnector.GetConnector().DB

	if err := DB.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) UpdateProduct(product *model.Product) error {
	DB := r.repoConnector.GetConnector().DB

	productToUpdate, err := r.GetProduct(product.Id)
	if err != nil {
		return err
	}

	if productToUpdate.Id < 1 {
		return fmt.Errorf("product with %d not found", product.Id)
	}

	productToUpdate.Owner = product.Owner
	productToUpdate.Price = product.Price
	productToUpdate.State = product.State

	if err := DB.Where("id = ?",
		product.Id).Update(&product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) CreateImage(productId uint, file *bufio.Reader) (model.Image, error) {
	image := model.Image{}

	DB := r.repoConnector.GetConnector().DB
	if _, err := r.getProductModel(productId); err != nil {
		return image, err
	}

	imageNameBase := fmt.Sprintf("%d-%d-%d", productId, time.Now().UnixNano(), rand.Int63())
	imageNameKeccak := crypto.Keccak256([]byte(imageNameBase))

	image.Name = hex.EncodeToString(imageNameKeccak)
	image.ProductId = uint(productId)

	fileContents, _ := ioutil.ReadAll(file)

	if err := os.WriteFile(config.IMAGE_LOCAL_STORAGE+image.Name, fileContents, 0644); err != nil {
		return image, err
	}

	if err := DB.Create(&image).Error; err != nil {
		return image, err
	}

	return image, nil
}

func (r *productRepository) GetImage(fileName string) (string, string, *bufio.Reader, error) {
	directory := config.IMAGE_LOCAL_STORAGE
	file, err := os.Open(directory + "/" + fileName)

	return fileName, directory, bufio.NewReader(file), err
}

func (r *productRepository) getProductModel(productId uint) (model.ProductModel, error) {
	DB := r.repoConnector.GetConnector().DB
	var product model.ProductModel

	if err := DB.Where("id = ?",
		productId).Preload("Img").Preload("Specification").First(&product, productId).Error; err != nil {
		return model.ProductModel{}, err
	}

	return product, nil
}
