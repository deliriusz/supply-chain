package main

import (
	"rafal-kalinowski.pl/adapter/api"
	"rafal-kalinowski.pl/config"
	"rafal-kalinowski.pl/domain/model"
	domain "rafal-kalinowski.pl/domain/service"
	"rafal-kalinowski.pl/repository"
)

func main() {
	config.Init()
	model.ConnectDatabase() //TODO: delete

	repoConnector := repository.NewRepoConnector()
	if err := repoConnector.InitConnection("firmex.db", ""); err != nil {
		panic(err)
	}

	loginRepository := repository.NewLoginRepository(repoConnector)
	loginService := domain.NewLoginService(loginRepository)
	productRepository := repository.NewProductRepository(repoConnector)
	productService := domain.NewProductService(productRepository)
	purchaseRepository := repository.NewPurchaseRepository(repoConnector)
	purchaseService := domain.NewPurchaseService(purchaseRepository)

	httpApi := api.NewHTTPHandler(loginService, productService, purchaseService)

	httpApi.Init()
	if err := httpApi.Start(); err != nil {
		panic(err)
	}
}
