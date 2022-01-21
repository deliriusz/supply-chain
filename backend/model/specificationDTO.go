package model

type SpecificationDTO struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value"`
}
