package model

type ImageDTO struct {
	Id        uint   `json:"id"`
	ProductId uint   `json:"productId"`
	Name      string `json:"name"`
	Url       string `json:"url"`
}
