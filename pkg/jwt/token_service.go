package jwt

import "github.com/golang-jwt/jwt/v5"

type tokenService struct {
	signSecretToken []byte
}

type TokenService interface {
	Create(claims map[string]interface{}) (string, error)
	Parse(token string) (map[string]interface{}, error)
}

func NewTokenService(signSecretToken string) TokenService {
	return &tokenService{
		signSecretToken: []byte(signSecretToken),
	}
}

func (t tokenService) Create(claims map[string]interface{}) (string, error) {
	var mapClaims jwt.MapClaims
	mapClaims = claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	tokenString, err := token.SignedString(t.signSecretToken)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t tokenService) Parse(token string) (map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}
