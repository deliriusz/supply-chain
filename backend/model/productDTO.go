package model

type ProductDTO struct {
	Id                uint               `json:"id" binding:"optional"`
	Img               []ImageDTO         `json:"images" binding:"optional"`
	Title             string             `json:"title" binding:"required"`
	Description       string             `json:"description" binding:"required"`
	Price             uint               `json:"price" binding:"required"`
	AvailableQuantity uint               `json:"availableQuantity" binding:"required"`
	Specification     []SpecificationDTO `json:"specification" binding:"required"`
}
