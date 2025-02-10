package sessions

import (
	"auth/infrastructure/redis/commands"
	"auth/internal/entities"
	"auth/internal/repositories"
	"context"
	"github.com/redis/go-redis/v9"
)

type selectSessionByAccessTokenCommand struct {
	client *redis.Client
}

func NewSelectSessionByAccessTokenCommand(client *redis.Client) repositories.SelectSessionByAccessTokenCommand {
	return &selectSessionByAccessTokenCommand{client: client}
}

func (s selectSessionByAccessTokenCommand) Execute(context context.Context, accessToken string) (entities.Session, error) {
	key := getKey(accessToken)
	value, err := commands.GetValueOrNil[entities.Session](context, s.client, key)
	if err != nil {
		return entities.Session{}, err
	}
	if value == nil {
		return entities.Session{}, repositories.ErrEntityNotFound
	}

	return *value, nil
}
