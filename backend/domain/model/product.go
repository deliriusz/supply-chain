package model

type Product struct {
	Id            uint            `json:"id" gorm:"primaryKey"`
	Img           []Image         `json:"img" gorm:"foreignKey:ProductId"`
	Title         string          `json:"title"`
	Description   string          `json:"description"`
	Price         uint            `json:"price"`
	Quantity      uint            `json:"quantity"`
	Specification []Specification `json:"specification" gorm:"foreignKey:ProductId"`
}
