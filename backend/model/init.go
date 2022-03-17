package model

import (
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "firmex.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Image{})
	database.AutoMigrate(&Specification{})
	database.AutoMigrate(&Product{})
	database.AutoMigrate(&PurchaseOrder{})

	DB = database
}

func Paginate(limit string, offset string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		limit, err := strconv.Atoi(limit)
		if err != nil {
			limit = 10
		} else if limit > 100 {
			limit = 100
		}

		offset, err := strconv.Atoi(offset)
		if err != nil {
			offset = 0
		} else if offset < 0 {
			offset = 0
		}

		return db.Offset(offset).Limit(limit)
	}
}
