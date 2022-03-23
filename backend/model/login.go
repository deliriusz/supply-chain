package model

type Login struct {
	Id        uint   `gorm:"primary_key"`
	Address   string ``
	SessionId string ``
	ExpiresAt int64  ``
}
