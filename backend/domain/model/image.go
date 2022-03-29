package model

type Image struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	ProductId uint   `json:"productId"`
	ImageName string `json:"imageName"`
}
