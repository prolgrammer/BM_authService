package app

import (
	"auth/infrastructure/postgres"
	"auth/infrastructure/postgres/commands/account"
	"auth/infrastructure/postgres/commands/session"
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

func CreateSessionRepository(client *postgres.Client) repositories.SessionRepository {
	insertSessionCommand := session.NewInsertSessionCommand(client)

	return repositories.NewSessionRepository(
		insertSessionCommand,
	)
}
