package account

import (
	"context"
	"github.com/prolgrammer/BM_authService/infrastructure/postgres"
	"github.com/prolgrammer/BM_authService/infrastructure/postgres/commands"
	"github.com/prolgrammer/BM_authService/internal/entities"
	"github.com/prolgrammer/BM_authService/internal/repositories"
)

type insertAccountPGCommand struct {
	client *postgres.Client
}

func NewInsertAccountCommand(client *postgres.Client) repositories.InsertAccountCommand {
	return &insertAccountPGCommand{client: client}
}

func (ic *insertAccountPGCommand) Execute(context context.Context, account entities.Account) (string, error) {
	sql, args, err := ic.client.Builder.Insert(commands.AccountTable).
		Columns(
			commands.AccountEmailField,
			commands.AccountPasswordField,
			commands.AccountRegistrationDateField,
			commands.AccountIsVerifiedField,
		).
		Values(account.Email, account.Password, account.RegistrationDate, false).
		Suffix("RETURNING " + commands.AccountIdField).
		ToSql()
	if err != nil {
		return "", err
	}

	var id string
	err = ic.client.Pool.QueryRow(context, sql, args...).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
