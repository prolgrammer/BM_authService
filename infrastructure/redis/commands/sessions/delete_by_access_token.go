package sessions

import (
	"auth/internal/repositories"
	"context"
	"github.com/redis/go-redis/v9"
)

type newDeleteByAccessTokenCommand struct {
	client *redis.Client
}

func NewDeleteByAccessTokenCommand(client *redis.Client) repositories.DeleteSessionByAccessTokenCommand {
	return &newDeleteByAccessTokenCommand{client: client}
}

func (n newDeleteByAccessTokenCommand) Execute(context context.Context, accessToken string) error {
	key := getKey(accessToken)
	err := n.client.Del(context, key).Err()

	return err
}
