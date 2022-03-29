package model

type PurchaseOrderDTO struct {
	Id      uint    `json:"id" gorm:"primary_key"`
	UserId  uint    `json:"userId"`
	Product Product `json:"product"`
	Price   uint    `json:"price"`
	Date    string  `json:"date"`
	Status  string  `json:"status"`
}
