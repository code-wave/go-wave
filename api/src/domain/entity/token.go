package entity

import "encoding/json"

type RefreshToken struct {
	Uuid         string `json:"uuid"`
	RefreshToken string `json:"refresh_token"`
	UserID       int64  `json:"user_id"`
	ExpiresAt    int64  `json:"expires_at"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

func (at *AccessToken) ResponseJSON() interface{} {
	atJson, _ := json.Marshal(at)
	return atJson
}
