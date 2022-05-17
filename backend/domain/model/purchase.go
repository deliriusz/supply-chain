package model

type PurchaseOrder struct {
	Id      uint      `json:"id" gorm:"primaryKey"`
	UserId  string    `json:"userId"`
	Product []Product `json:"product" gorm:"foreignKey:Id"`
	Date    string    `json:"date"`
}
