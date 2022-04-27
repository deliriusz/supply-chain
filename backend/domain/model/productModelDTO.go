package model

type ProductModelDTO struct {
	Id            uint               `json:"id"`
	Img           []ImageDTO         `json:"images"`
	Title         string             `json:"title" binding:"required"`
	Description   string             `json:"description" binding:"required"`
	Price         uint               `json:"price" binding:"required"`
	Quantity      uint               `json:"quantity" binding:"required"`
	Specification []SpecificationDTO `json:"specification" binding:"required"`
}
