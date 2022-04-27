package model

type LoginDTO struct {
	Address   string `json:"address"`
	Role      string `json:"role"`
	ExpiresAt int64  `json:"expiresAt,omitempty"`
	TTL       uint   `json:"ttl,omitempty"`
}
