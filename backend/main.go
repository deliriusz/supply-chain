package main

import (
	"rafal-kalinowski.pl/adapter/api"
	"rafal-kalinowski.pl/config"
	domain "rafal-kalinowski.pl/domain/service"
	"rafal-kalinowski.pl/repository"
)

func main() {
	config.Init()

	dbRepoConnector := repository.GetProvider[*repository.DBRepoConnector](repository.ProviderFactory)
	if err := dbRepoConnector.InitConnection("firmex.db", ""); err != nil {
		panic(err)
	}

	ethRepoConnector := repository.GetProvider[*repository.EthereumRepoConnector](repository.ProviderFactory)
	if err := ethRepoConnector.InitConnection("firmex.db", ""); err != nil {
		panic(err)
	}

	loginRepository := repository.NewLoginRepository(dbRepoConnector, ethRepoConnector)
	loginService := domain.NewLoginService(loginRepository)
	productRepository := repository.NewProductRepository(dbRepoConnector)
	productService := domain.NewProductService(productRepository)
	purchaseRepository := repository.NewPurchaseRepository(dbRepoConnector)
	purchaseService := domain.NewPurchaseService(purchaseRepository)

	httpApi := api.NewHTTPHandler(loginService, productService, purchaseService)

	httpApi.Init()
	if err := httpApi.Start(); err != nil {
		panic(err)
	}
}
