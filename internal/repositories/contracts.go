package repositories

import (
	"auth/internal/entities"
	"context"
)

type (
	InsertAccountCommand interface {
		Execute(context context.Context, account entities.Account) (string, error)
	}

	SelectAccountByEmailCommand interface {
		Execute(context context.Context, email entities.Email) (entities.Account, error)
	}
)

type (
	InsertSessionCommand interface {
		Execute(context context.Context, account entities.Session) error
	}

	SelectSessionByAccessTokenCommand interface {
		Execute(context context.Context, accessToken string) (entities.Session, error)
	}

	UpdateSessionByAccessTokenCommand interface {
		Execute(context context.Context, accessToken string, newSession entities.Session) error
	}

	DeleteSessionByAccessTokenCommand interface {
		Execute(context context.Context, accessToken string) error
	}
)
