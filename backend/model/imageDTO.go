package model

type ImageDTO struct {
	Id        uint   `json:"id" binding:"optional"`
	ProductId uint   `json:"productId" binding:"optional"`
	Url       string `json:"url" binding:"required"`
	ImageName string `json:"imageName" binding:"optional"`
}
