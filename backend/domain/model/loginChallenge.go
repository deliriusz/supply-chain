package model

type LoginChallenge struct {
	Signature string `json:"signature,omitempty"`
	Data      string `json:"data,omitempty"`
	Address   string `json:"address,omitempty"`
	Nonce     int64  `json:"nonce,string,omitempty"`
}
