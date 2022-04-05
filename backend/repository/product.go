package repository

import (
	"bufio"
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
	repoConnector RepoConnector
}

func NewProductRepository(c RepoConnector) domain.ProductRepository {
	repo := &productRepository{
		repoConnector: c,
	}

	return repo
}

func (r *productRepository) GetProducts(limit, offset uint) ([]model.Product, uint) {
	DB := r.repoConnector.GetConnector()
	var products []model.Product

	count := uint(0)
	DB.Find(&[]model.Product{}).Count(&count)

	DB.Preload("Img").Preload("Specification").
		Scopes(Paginate(int(limit), int(offset))).
		Find(&products)

	return products, count
}

func (r *productRepository) GetProduct(id uint) (model.Product, error) {
	product, err := r.getProduct(id)

	return product, err
}

func (r *productRepository) CreateProduct(product *model.Product) error {
	DB := r.repoConnector.GetConnector()
	if err := DB.Create(&product).Error; err != nil {
		return err
	}

	for _, sd := range product.Specification {
		DB.Create(&sd)
	}

	return nil
}

func (r *productRepository) CreateImage(productId uint, file *bufio.Reader) (string, error) {
	DB := r.repoConnector.GetConnector()
	if _, err := r.getProduct(productId); err != nil {
		return "", err
	}

	imageNameBase := fmt.Sprintf("%d-%d-%d", productId, time.Now().UnixNano(), rand.Int63())
	imageNameKeccak := crypto.Keccak256([]byte(imageNameBase))

	image := model.Image{
		ProductId: uint(productId),
		ImageName: hex.EncodeToString(imageNameKeccak),
	}

	fileContents, _ := ioutil.ReadAll(file)

	if err := os.WriteFile(config.IMAGE_LOCAL_STORAGE+image.ImageName, fileContents, 0644); err != nil {
		return "", err
	}

	if err := DB.Create(&image).Error; err != nil {
		return "", err
	}

	return image.ImageName, nil
}

func (r *productRepository) GetImage(fileName string) (string, string, *bufio.Reader, error) {
	directory := config.IMAGE_LOCAL_STORAGE
	file, err := os.Open(directory + "/" + fileName)

	return fileName, directory, bufio.NewReader(file), err
}

func (r *productRepository) getProduct(productId uint) (model.Product, error) {
	DB := r.repoConnector.GetConnector()
	var product model.Product

	if err := DB.Where("id = ?",
		productId).Preload("Img").Preload("Specification").First(&product, productId).Error; err != nil {
		return model.Product{}, err
	}

	return product, nil
}
