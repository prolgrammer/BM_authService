package session

import (
	"auth/infrastructure/postgres"
	"auth/infrastructure/postgres/commands"
	"auth/internal/entities"
	"auth/internal/repositories"
	"context"
)

type insertSessionPGCommand struct {
	client *postgres.Client
}

func NewInsertSessionCommand(client *postgres.Client) repositories.InsertSessionCommand {
	return &insertSessionPGCommand{
		client: client,
	}
}

func (i *insertSessionPGCommand) Execute(context context.Context, session entities.Session) error {
	sql, args, err := i.client.Builder.
		Insert(commands.SessionTable).
		Columns(
			commands.SessionAccessTokenField,
			commands.SessionRefreshTokenField,
			commands.SessionAccountIdField,
			commands.SessionExpiresAtField).
		Values(
			session.AccessToken,
			session.RefreshToken,
			session.AccountId,
			session.ExpiresAt).
		Suffix("RETURNING id").ToSql()
	if err != nil {
		return err
	}

	var id string
	err = i.client.Pool.QueryRow(context, sql, args...).Scan(&id)

	return err
}
