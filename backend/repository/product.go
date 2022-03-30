package repository

import (
	"bytes"
	"fmt"
	"math/rand"
	"mime/multipart"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"rafal-kalinowski.pl/config"
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
	var products []model.Product

	count := 0
	model.DB.Find(&[]model.Product{}).Count(&count)

	model.DB.Preload("Img").Preload("Specification").
		Scopes(Paginate(int(limit), int(offset))).
		Find(&products)

	return products
}

func (repo *productRepository) GetProduct(id uint) (model.Product, error) {
	product, err := getProduct(id)

	return product, err
}

func (repo *productRepository) CreateProduct(product *model.Product) error {
	if err := model.DB.Create(&product).Error; err != nil {
		return err
	}

	for _, sd := range product.Specification {
		model.DB.Create(&sd)
	}

	return nil
}

func (repo *productRepository) CreateImage(productId uint, file *multipart.File) error {
	if _, err := getProduct(productId); err != nil {
		return err
	}

	imageNameBase := fmt.Sprintf("%d-%d-%d", productId, time.Now().UnixNano(), rand.Int63())
	imageNameKeccak := crypto.Keccak256([]byte(imageNameBase))

	image := model.Image{
		ProductId: uint(productId),
		ImageName: string(imageNameKeccak),
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(*file)

	// err = nil
	if err := os.WriteFile(config.IMAGE_LOCAL_STORAGE+image.ImageName, buf.Bytes(), 0644); err != nil {
		return err
	}

	if err := model.DB.Create(&image).Error; err != nil {
		return err
	}

	return nil
}

func (repo *productRepository) GetImage(fileName string) (string, string, *os.File, error) {
	directory := config.IMAGE_LOCAL_STORAGE
	file, err := os.Open(directory + "/" + fileName)

	return fileName, directory, file, err
}

func getProduct(productId uint) (model.Product, error) {
	var product model.Product

	if err := model.DB.Where("id = ?",
		productId).Preload("Img").Preload("Specification").First(&product, productId).Error; err != nil {
		return model.Product{}, err
	}

	return product, nil
}
