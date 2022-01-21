package model

type Specification struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	ProductId uint   `json:"productId"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}
