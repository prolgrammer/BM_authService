package usecases

import (
	"context"
	"github.com/prolgrammer/BM_authService/internal/entities"
)

type (
	SignUpAccountRepository interface {
		CheckEmailExists(ctx context.Context, email entities.Email) (bool, error)
		Insert(ctx context.Context, email entities.Account) (string, error)
	}

	SignInAccountRepository interface {
		SelectByEmail(ctx context.Context, email string) (entities.Account, error)
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

type (
	SignUpHashService interface {
		CreateHash(password string) ([]byte, error)
	}

	SignInHashService interface {
		CompareStringAndHash(stringToCompare string, hashedString string) (bool, error)
	}
)
