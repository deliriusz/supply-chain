package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"rafal-kalinowski.pl/domain/model"
)

type RepoConnector interface {
	InitConnection(name, url string) error
	GetConnector() *gorm.DB
}

type connector struct {
	DB *gorm.DB
}

// GetConnector implements RepoConnector
func (c *connector) GetConnector() *gorm.DB {
	return c.DB
}

// InitConnection implements RepoConnector
func (c *connector) InitConnection(name, url string) error {
	database, err := gorm.Open("sqlite3", name)

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&model.Image{})
	database.AutoMigrate(&model.Specification{})
	database.AutoMigrate(&model.Product{})
	database.AutoMigrate(&model.PurchaseOrder{})
	database.AutoMigrate(&model.Login{})

	c.DB = database

	return err
}

func NewRepoConnector() RepoConnector {
	repo := &connector{
		DB: nil,
	}

	return repo
}

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
