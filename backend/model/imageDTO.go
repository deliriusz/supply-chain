package model

type ImageDTO struct {
	Id        uint   `json:"id"`
	ProductId uint   `json:"productId"`
	Url       string `json:"url" binding:"required"`
}
