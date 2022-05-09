package model

type ProductModelDTO struct {
	Id            uint               `json:"id"`
	Img           []ImageDTO         `json:"images"`
	Title         string             `json:"title"`
	Description   string             `json:"description"`
	BasePrice     uint               `json:"price"`
	Quantity      uint               `json:"quantity"`
	Specification []SpecificationDTO `json:"specification"`
}
