package app

import (
	"auth/infrastructure/postgres"
	"auth/infrastructure/postgres/commands/account"
	"auth/infrastructure/redis/commands/sessions"
	"auth/internal/repositories"
	"github.com/redis/go-redis/v9"
)

func CreatePGAccountRepository(client *postgres.Client) repositories.AccountRepository {
	insertAccountCommand := account.NewInsertAccountCommand(client)
	selectAccountByEmailCommand := account.NewSelectAccountByEmail(client)

	return repositories.NewAccountRepository(
		insertAccountCommand,
		selectAccountByEmailCommand,
	)
}

func CreateSessionRepository(client *redis.Client) repositories.SessionRepository {
	insertSessionCommand := sessions.NewInsertSessionRedisCommand(client)
	selectSessionByAccessTokenCommand := sessions.NewSelectSessionByAccessTokenCommand(client)
	updateSessionByAccessTokenCommand := sessions.NewUpdateByAccessTokenCommand(client)
	deleteSessionByAccessTokenCommand := sessions.NewDeleteByAccessTokenCommand(client)

	return repositories.NewSessionRepository(
		insertSessionCommand,
		selectSessionByAccessTokenCommand,
		updateSessionByAccessTokenCommand,
		deleteSessionByAccessTokenCommand,
	)
}
