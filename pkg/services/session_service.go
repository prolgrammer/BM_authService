package pkg

import (
	"auth/config"
	"auth/internal/entities"
	"auth/pkg/jwt"
	"time"
)

type sessionService struct {
	config       config.TokenConfig
	accessToken  jwt.TokenService
	refreshToken jwt.TokenService
}

type SessionService interface {
	CreateSession(account entities.Account) (entities.Session, error)
}

func NewSessionService(config config.TokenConfig, accessToken jwt.TokenService, refreshToken jwt.TokenService) SessionService {
	return &sessionService{
		config:       config,
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}
}

func (s *sessionService) CreateSession(account entities.Account) (entities.Session, error) {
	accessExpiresAt := time.Now().Add(s.config.AccessTokenDuration)
	refreshExpiresAt := time.Now().Add(s.config.RefreshTokenDuration)

	accessClaims := entities.NewClaims(account.Id, account.Role, accessExpiresAt)
	access, err := s.accessToken.Create(accessClaims)
	if err != nil {
		return entities.Session{}, err
	}

	refreshClaims := entities.NewClaims(account.Id, account.Role, refreshExpiresAt)
	refresh, err := s.refreshToken.Create(refreshClaims)
	if err != nil {
		return entities.Session{}, err
	}

	return entities.Session{
		AccountId:            account.Id,
		AccessToken:          access,
		AccessTokenExpiresAt: accessExpiresAt,
		RefreshToken:         refresh,
		ExpiresAt:            refreshExpiresAt,
	}, nil
}
