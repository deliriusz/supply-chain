package model

import "time"

type PurchaseOrder struct {
	Id      uint      `json:"id" gorm:"primaryKey"`
	UserId  string    `json:"userId"`
	Product []Product `json:"product" gorm:"foreignKey:Id"`
	Date    time.Time `json:"date"`
}
