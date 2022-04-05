package repository_test

import (
	"os"

	"rafal-kalinowski.pl/config"
	domain "rafal-kalinowski.pl/domain/service"
	"rafal-kalinowski.pl/repository"
)

var TABLE_NAME string
var loginRepo domain.LoginRepository
var productRepo domain.ProductRepository
var purchaseRepo domain.PurchaseRepository

func init() {
	TABLE_NAME = "firmex-repo-test.db"
	config.Init("../.env")
}

func Setup() {
	repoConn := repository.NewRepoConnector()
	if err := repoConn.InitConnection(TABLE_NAME, ""); err != nil {
		panic(err)
	}

	productRepo = repository.NewProductRepository(repoConn)
	loginRepo = repository.NewLoginRepository(repoConn)
	purchaseRepo = repository.NewPurchaseRepository(repoConn)
}

func Cleanup() {
	os.Remove(TABLE_NAME)
}