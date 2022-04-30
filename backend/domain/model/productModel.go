package model

// ProductModel is just a set of characteristins for a model of a product.
// Unique products for a model are then created and NFT is generated for them.
type ProductModel struct {
	Id            uint            `gorm:"primaryKey"`
	Img           []Image         `gorm:"foreignKey:ProductId"`
	Title         string          ``
	Description   string          ``
	BasePrice     uint            ``
	Quantity      uint            ``
	Specification []Specification `gorm:"foreignKey:ProductId"`
}
