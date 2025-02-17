package app

import (
	"github.com/prolgrammer/BM_authService/infrastructure/postgres"
	"github.com/prolgrammer/BM_authService/infrastructure/postgres/commands/account"
	"github.com/prolgrammer/BM_authService/infrastructure/redis/commands/sessions"
	"github.com/prolgrammer/BM_authService/internal/repositories"
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
