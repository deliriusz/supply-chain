package model

type ProductDTO struct {
	Id    uint            `json:"id"`
	Model ProductModelDTO `json:"productModel"`
	State string          `json:"state"`
	Owner string          `json:"owner,omitempty"`
	Price uint            `json:"price"`
}
