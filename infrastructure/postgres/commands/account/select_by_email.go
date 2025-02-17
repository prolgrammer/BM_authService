package account

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/prolgrammer/BM_authService/infrastructure/postgres"
	"github.com/prolgrammer/BM_authService/infrastructure/postgres/commands"
	"github.com/prolgrammer/BM_authService/internal/entities"
	"github.com/prolgrammer/BM_authService/internal/repositories"
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
