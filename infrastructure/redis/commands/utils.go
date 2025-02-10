package commands

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

func SetValue(ctx context.Context, client *redis.Client, key string, value any, expiration time.Duration) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = client.Set(ctx, key, json, expiration).Err()
	return err
}

func GetValueOrNil[V any](ctx context.Context, client *redis.Client, key string) (*V, error) {
	str, err := GetStringValueIfExist(ctx, client, key)
	if err != nil {
		return nil, err
	}
	if str == "" {
		return nil, nil
	}

	var value V
	err = json.Unmarshal([]byte(str), &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func GetStringValueIfExist(ctx context.Context, client *redis.Client, key string) (string, error) {
	cmd := client.Get(ctx, key)

	if cmd.Err() != nil {
		if !errors.Is(cmd.Err(), redis.Nil) {
			return "", cmd.Err()
		}
		return "", nil
	}

	return cmd.Result()
}
