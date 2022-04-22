package repository

import (
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"rafal-kalinowski.pl/domain/model"
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

type DBRepoConnector struct {
	DB *gorm.DB
}

// GetConnector implements RepoConnector
func (c *DBRepoConnector) GetConnector() *DBRepoConnector {
	return c
}

// InitConnection implements RepoConnector
func (c *DBRepoConnector) InitConnection(name, connectionString string) error {
	database, err := gorm.Open("sqlite3", name)

	if err != nil {
		panic(err)
	}

	x := EthereumRepoConnector{}
	x.GetConnector()

	database.AutoMigrate(&model.Image{})
	database.AutoMigrate(&model.Specification{})
	database.AutoMigrate(&model.Product{})
	database.AutoMigrate(&model.PurchaseOrder{})
	database.AutoMigrate(&model.Login{})

	c.DB = database

	return err
}

// func NewRepoConnector[T RepoType](v T) RepoConnector[T] {
func NewRepoConnector[T RepoType](v T) RepoConnector[*EthereumRepoConnector] {
	x := EthereumRepoConnector{}
	return &x
	// return &v
}

// func NewRepoConnector[T RepoType](v T) RepoConnector[T] {
// 	return v
// }

func Paginate(limit, offset int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if limit <= 0 {
			limit = 10
		} else if limit > 100 {
			limit = 100
		}

		if offset < 0 {
			offset = 0
		}

		return db.Offset(offset).Limit(limit)
	}
}

func init() {
	ProviderFactory = &DataProviderFactory{
		dataProviders: make(map[string]any),
	}

	RegisterProvider[*EthereumRepoConnector](ProviderFactory, &EthereumRepoConnector{})
	RegisterProvider[*DBRepoConnector](ProviderFactory, &DBRepoConnector{})
}
