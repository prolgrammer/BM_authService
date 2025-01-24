package entities

import "time"

type Session struct {
	AccountId            string
	AccessToken          string
	RefreshToken         string
	AccessTokenExpiresAt time.Time
	ExpiresAt            time.Time
}
