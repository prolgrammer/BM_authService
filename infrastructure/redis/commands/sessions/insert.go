package sessions

import (
	"context"
	"github.com/prolgrammer/BM_authService/infrastructure/redis/commands"
	"github.com/prolgrammer/BM_authService/internal/entities"
	"github.com/prolgrammer/BM_authService/internal/repositories"
	"github.com/redis/go-redis/v9"
	"time"
)

type insertSessionRedisCommand struct {
	client *redis.Client
}

func NewInsertSessionRedisCommand(client *redis.Client) repositories.InsertSessionCommand {
	return &insertSessionRedisCommand{client: client}
}

func (i *insertSessionRedisCommand) Execute(context context.Context, session entities.Session) error {
	key := getKey(session.AccessToken)
	sessionPtr, err := commands.GetValueOrNil[entities.Session](context, i.client, key)
	if err != nil {
		return err
	}
	if sessionPtr != nil {
		return err
	}

	err = commands.SetValue(context, i.client, key, session, session.ExpiresAt.Sub(time.Now()))
	return err
}
