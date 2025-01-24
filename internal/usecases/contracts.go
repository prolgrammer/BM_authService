package usecases

import (
	"auth/internal/entities"
	"context"
)

type (
	SignUpAccountRepository interface {
		CheckEmailExists(ctx context.Context, email entities.Email) (bool, error)
		Insert(ctx context.Context, email entities.Account) (string, error)
	}
)

type (
	SessionService interface {
		CreateSession(user entities.Account) (entities.Session, error)
	}

	SessionRepository interface {
		Insert(ctx context.Context, session entities.Session) error
	}
)
