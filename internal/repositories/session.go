package repositories

import (
	"auth/internal/entities"
	"context"
)

type sessionRepository struct {
	insertSessionCommand InsertSessionCommand
}

type SessionRepository interface {
	Insert(ctx context.Context, session entities.Session) error
}

func NewSessionRepository(insertSessionCommand InsertSessionCommand) SessionRepository {
	return &sessionRepository{
		insertSessionCommand: insertSessionCommand,
	}
}

func (s sessionRepository) Insert(ctx context.Context, session entities.Session) error {
	return s.insertSessionCommand.Execute(ctx, session)
}
