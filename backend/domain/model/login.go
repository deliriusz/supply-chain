package model

type Login struct {
	Id        uint     `gorm:"primaryKey"`
	Address   string   ``
	Role      UserRole ``
	SessionId string   ``
	ExpiresAt int64    ``
	TTL       uint     ``
}
