package model

type Login struct {
	Id        uint   `gorm:"primaryKey"`
	Address   string ``
	SessionId string ``
	ExpiresAt int64  ``
	TTL       uint   ``
}
