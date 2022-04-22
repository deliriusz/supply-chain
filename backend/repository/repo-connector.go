package repository

import (
	"reflect"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var ProviderFactory *DataProviderFactory

type DataProviderFactory struct {
	dataProviders map[string]any
}

func GetProvider[T RepoType](factory *DataProviderFactory) RepoConnector[T] {
	var provider T
	prov, has := factory.dataProviders[reflect.TypeOf(provider).Name()]
	if !has {
		return nil
	}

	return prov.(RepoConnector[T])
}

func RegisterProvider[T RepoType](factory *DataProviderFactory, p RepoConnector[T]) {
	var provider T
	factory.dataProviders[reflect.TypeOf(provider).Name()] = p
}

type RepoType interface {
	*EthereumRepoConnector | *DBRepoConnector
}

type RepoConnector[T RepoType] interface {
	InitConnection(name, url string) error
	GetConnector() T
}

// GetConnector implements RepoConnector

// func NewRepoConnector[T RepoType](v T) RepoConnector[T] {
func NewRepoConnector[T RepoType](v T) RepoConnector[*EthereumRepoConnector] {
	x := EthereumRepoConnector{}
	return &x
}

func init() {
	ProviderFactory = &DataProviderFactory{
		dataProviders: make(map[string]any),
	}

	RegisterProvider[*EthereumRepoConnector](ProviderFactory, &EthereumRepoConnector{})
	RegisterProvider[*DBRepoConnector](ProviderFactory, &DBRepoConnector{})
}
