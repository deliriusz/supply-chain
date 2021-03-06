package repository_test

import (
	"io/ioutil"
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
	dbRepoConn := repository.GetProvider[*repository.DBRepoConnector](repository.ProviderFactory)
	ethRepoConn := repository.GetProvider[*repository.EthereumRepoConnector](repository.ProviderFactory)

	if err := dbRepoConn.InitConnection(TABLE_NAME, ""); err != nil {
		panic(err)
	}

	productRepo = repository.NewProductRepository(dbRepoConn)
	loginRepo = repository.NewLoginRepository(dbRepoConn, ethRepoConn)
	purchaseRepo = repository.NewPurchaseRepository(dbRepoConn)
}

func Cleanup() {
	os.Remove("./" + TABLE_NAME)
	dir, _ := ioutil.ReadDir(config.IMAGE_LOCAL_STORAGE)
	for _, d := range dir {
		os.RemoveAll(config.IMAGE_LOCAL_STORAGE + d.Name())
	}
}
