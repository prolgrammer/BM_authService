package account

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/prolgrammer/BM_authService/infrastructure/postgres"
	"github.com/prolgrammer/BM_authService/internal/entities"
	"github.com/prolgrammer/BM_authService/internal/repositories"
)

func selectAccount(context context.Context, client *postgres.Client, sql string, args []any) (entities.Account, error) {
	result := entities.Account{}
	row := client.Pool.QueryRow(context, sql, args...)
	err := row.Scan(&result.Id, &result.Email, &result.Password, &result.RegistrationDate, &result.Role)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entities.Account{}, repositories.ErrEntityNotFound
		}
		return entities.Account{}, err
	}
	return result, nil
}
