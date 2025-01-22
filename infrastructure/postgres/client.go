package postgres

import (
	"auth/config/pg"
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = 10 * time.Second
)

var (
	ErrNoChange = errors.New("no changes applied")
)

type Client struct {
	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
	cfg     pg.Config
}

func NewClient(cfg pg.Config) (*Client, error) {
	client := &Client{cfg: cfg}

	connAttempts := cfg.RetryConnectionAttempts
	connTimeout := cfg.RetryConnectionTimeout
	maxPoolSize := cfg.MaxPoolSize

	if maxPoolSize == 0 {
		maxPoolSize = _defaultMaxPoolSize
	}

	if connAttempts == 0 {
		connAttempts = _defaultConnAttempts
	}

	if connTimeout == 0 {
		connTimeout = _defaultConnTimeout
	}

	client.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	connectionString := client.connectionString()

	poolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		fmt.Println("couldn't parse connection string")
		return nil, err
	}

	poolConfig.MaxConns = int32(maxPoolSize)
	for connAttempts > 0 {
		client.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err == nil {
			err = client.Pool.Ping(context.TODO())
			if err == nil {
				break
			}
		}

		fmt.Println("failed to connect to postgres")
		<-time.After(connTimeout)

		connAttempts--
	}

	if err != nil {
		fmt.Println("couldn't connect to postgres")
		return nil, err
	}

	return client, nil
}

func (c *Client) connectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.cfg.User,
		c.cfg.Password,
		c.cfg.Host,
		c.cfg.Port,
		c.cfg.Database,
		c.cfg.SSLMode,
	)
}

func (c *Client) MigrateUp() error {
	m, err := migrate.New(c.cfg.MigrationsPath, c.connectionString())
	if err != nil {
		return fmt.Errorf("failed to create migration handler: %v", err)
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return ErrNoChange
		}
		return fmt.Errorf("failed to migrate up: %v", err)
	}

	return nil
}

func (c *Client) Close() {
	if c.Pool != nil {
		c.Pool.Close()
	}
}
