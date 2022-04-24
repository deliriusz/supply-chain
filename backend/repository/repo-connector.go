package repository

import (
	"reflect"
)

var ProviderFactory *DataProviderFactory

type DataProviderFactory struct {
	dataProviders map[string]any
}

func GetProvider[T RepoType](factory *DataProviderFactory) RepoConnector[T] {
	var provider T
	providerType := reflect.TypeOf(&provider).Elem()
	prov, has := factory.dataProviders[providerType.String()]
	if !has {
		return nil
	}

	return prov.(RepoConnector[T])
}

func RegisterProvider[T RepoType](factory *DataProviderFactory, p RepoConnector[T]) {
	var provider T
	providerType := reflect.TypeOf(&provider).Elem()
	factory.dataProviders[providerType.String()] = p
}

type RepoType interface {
	*EthereumRepoConnector | *DBRepoConnector
}

type RepoConnector[T RepoType] interface {
	InitConnection(name, url string) error
	GetConnector() T
}

func init() {
	ProviderFactory = &DataProviderFactory{
		dataProviders: make(map[string]any),
	}

	RegisterProvider[*EthereumRepoConnector](ProviderFactory, &EthereumRepoConnector{})
	RegisterProvider[*DBRepoConnector](ProviderFactory, &DBRepoConnector{})
}
