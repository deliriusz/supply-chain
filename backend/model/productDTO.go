package model

type ProductDTO struct {
	Id                uint               `json:"id"`
	Img               []ImageDTO         `json:"images"`
	Title             string             `json:"title" binding:"required"`
	Description       string             `json:"description" binding:"required"`
	Price             uint               `json:"price" binding:"required"`
	AvailableQuantity uint               `json:"quantity" binding:"required"`
	Specification     []SpecificationDTO `json:"specification" binding:"required"`
}
