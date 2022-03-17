package model

type PurchaseOrder struct {
	Id      uint    `json:"id" gorm:"primary_key"`
	UserId  uint    `json:"userId"`
	Product Product `json:"product" gorm:"foreignKey:Id"`
	Price   uint    `json:"price"`
	Date    string  `json:"date"`
	Status  string  `json:"status"`
}
