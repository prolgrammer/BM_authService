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
