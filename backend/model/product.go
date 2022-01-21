package model

type Product struct {
	Id                uint            `json:"id" gorm:"primary_key"`
	Img               []Image         `json:"img" gorm:"foreignKey:ProductId"`
	Title             string          `json:"title"`
	Description       string          `json:"description"`
	Price             uint            `json:"price"`
	AvailableQuantity uint            `json:"availableQuantity"`
	Specification     []Specification `json:"specification" gorm:"foreignKey:ProductId"`
}
