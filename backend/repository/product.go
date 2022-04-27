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

func (r *productRepository) GetProducts(limit, offset uint) ([]model.ProductModel, uint) {
	DB := r.repoConnector.GetConnector().DB
	var products []model.ProductModel

	count := uint(0)
	DB.Find(&[]model.ProductModel{}).Count(&count)

	DB.Preload("Img").Preload("Specification").
		Scopes(Paginate(int(limit), int(offset))).
		Find(&products)

	return products, count
}

func (r *productRepository) GetProduct(id uint) (model.ProductModel, error) {
	product, err := r.getProduct(id)

	return product, err
}

func (r *productRepository) CreateProduct(product *model.ProductModel) error {
	DB := r.repoConnector.GetConnector().DB
	if err := DB.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) CreateImage(productId uint, file *bufio.Reader) (model.Image, error) {
	image := model.Image{}

	DB := r.repoConnector.GetConnector().DB
	if _, err := r.getProduct(productId); err != nil {
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

func (r *productRepository) getProduct(productId uint) (model.ProductModel, error) {
	DB := r.repoConnector.GetConnector().DB
	var product model.ProductModel

	if err := DB.Where("id = ?",
		productId).Preload("Img").Preload("Specification").First(&product, productId).Error; err != nil {
		return model.ProductModel{}, err
	}

	return product, nil
}
