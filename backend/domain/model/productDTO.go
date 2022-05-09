package model

type ProductDTO struct {
	Id      uint            `json:"id"`
	ModelId uint            `json:"productModelId,omitempty"`
	Model   ProductModelDTO `json:"productModel,omitempty"`
	State   string          `json:"state"`
	Owner   string          `json:"owner,omitempty"`
	Price   uint            `json:"price"`
}
