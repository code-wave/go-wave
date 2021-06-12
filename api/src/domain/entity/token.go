package entity

import "encoding/json"

type RefreshToken struct {
	Uuid         string `json:"uuid"`
	RefreshToken string `json:"refresh_token"`
	UserID       uint64 `json:"user_id"`
	ExpiresAt    string `json:"expires_at"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
}

func (at *AccessToken) ResponseJSON() interface{} {
	aJson, _ := json.Marshal(at)
	return aJson
}
