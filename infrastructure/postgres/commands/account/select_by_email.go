package account

import (
	"auth/infrastructure/postgres"
	"auth/infrastructure/postgres/commands"
	"auth/internal/entities"
	"auth/internal/repositories"
	"context"
	"github.com/Masterminds/squirrel"
)

type selectAccountByEmail struct {
	client *postgres.Client
}

func NewSelectAccountByEmail(client *postgres.Client) repositories.SelectAccountByEmailCommand {
	return &selectAccountByEmail{client: client}
}

func (s selectAccountByEmail) Execute(context context.Context, email entities.Email) (entities.Account, error) {
	sql, args, err := s.client.Builder.
		Select(
			commands.AccountIdField,
			commands.AccountEmailField,
			commands.AccountPasswordField,
			commands.AccountRegistrationDateField,
			commands.AccountRoleField,
		).
		From(commands.AccountTable).
		Where(squirrel.Eq{commands.AccountEmailField: email}).
		ToSql()

	if err != nil {
		return entities.Account{}, err
	}

	return selectAccount(context, s.client, sql, args)
}
