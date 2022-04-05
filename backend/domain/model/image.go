package model

type Image struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	ProductId uint   `json:"productId"`
	ImageName string `json:"imageName"`
}
