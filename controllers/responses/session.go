package responses

type Session struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkzMjgyMjAsInJvbGUiOiJVU0VSIiwic3ViIjoiMyJ9.mp0uoVP-RTwOQrekQZm3PkjVnzdvUGfgnbYnT9piwaw"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDE4ODQyMjAsInJvbGUiOiJVU0VSIiwic3ViIjoiMyJ9.5ew-TEJ3io9kfxGQdO9F5b1KvPBW3REkutEAU9HypMQ"`
	ExpiresAt    int64  `json:"expires_at" example:"1741884220"`
}

func NewSession(
	accessToken string,
	refreshToken string,
	expiresAt int64) Session {
	return Session{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}

}
