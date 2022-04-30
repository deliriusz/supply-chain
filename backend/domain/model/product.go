package model

// Product represents a unique product created in a factory, that has an NFT assigned
type Product struct {
	Id    uint         `gorm:"primaryKey"`
	Model ProductModel `gorm:"foreignKey:Id"`
	State ProductState
	Owner string
	Price uint
}
