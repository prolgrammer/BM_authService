package sessions

import (
	"context"
	"github.com/prolgrammer/BM_authService/infrastructure/redis/commands"
	"github.com/prolgrammer/BM_authService/internal/entities"
	"github.com/prolgrammer/BM_authService/internal/repositories"
	"github.com/redis/go-redis/v9"
	"time"
)

type updateByAccessTokenCommand struct {
	client *redis.Client
}

func NewUpdateByAccessTokenCommand(client *redis.Client) repositories.UpdateSessionByAccessTokenCommand {
	return &updateByAccessTokenCommand{client: client}
}

func (u updateByAccessTokenCommand) Execute(context context.Context, accessToken string, newSession entities.Session) error {
	oldKey := getKey(accessToken)
	value, err := commands.GetValueOrNil[entities.Session](context, u.client, oldKey)
	if err != nil {
		return err
	}
	if value != nil {
		err = u.client.Del(context, oldKey).Err()
		if err != nil {
			return err
		}
	}

	newKey := getKey(newSession.AccessToken)

	err = commands.SetValue(context, u.client, newKey, newSession, newSession.ExpiresAt.Sub(time.Now()))
	if err != nil {
		return err
	}

	return nil
}
