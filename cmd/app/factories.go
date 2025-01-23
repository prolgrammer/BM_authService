package app

import (
	"auth/infrastructure/postgres"
	"auth/infrastructure/postgres/commands/account"
	"auth/internal/repositories"
)

func CreatePGAccountRepository(client *postgres.Client) repositories.AccountRepository {
	insertAccountCommand := account.NewInsertAccountCommand(client)
	selectAccountByEmailCommand := account.NewSelectAccountByEmail(client)

	return repositories.NewAccountRepository(
		insertAccountCommand,
		selectAccountByEmailCommand,
	)
}
