package model

// Product represents a unique product created in a factory, that has an NFT assigned
type Product struct {
	Id uint `gorm:"primaryKey"`
	// we use both ModelId and Model to indicate "belongs to" relation
	// otherwise gorm would create productModel together with product
	ModelId int
	Model   ProductModel `gorm:"foreignKey:ModelId"`
	State   ProductState
	Owner   string
	Price   uint
}
