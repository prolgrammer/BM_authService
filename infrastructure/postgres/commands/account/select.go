package account

import (
	"auth/infrastructure/postgres"
	"auth/internal/entities"
	"auth/internal/repositories"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
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
