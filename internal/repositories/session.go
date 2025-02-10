package repositories

import (
	"auth/internal/entities"
	"context"
)

type sessionRepository struct {
	insertSessionCommand              InsertSessionCommand
	selectSessionByAccessTokenCommand SelectSessionByAccessTokenCommand
	updateSessionByAccessTokenCommand UpdateSessionByAccessTokenCommand
	deleteSessionByAccessTokenCommand DeleteSessionByAccessTokenCommand
}

type SessionRepository interface {
	Insert(ctx context.Context, session entities.Session) error
}

func NewSessionRepository(
	insertSessionCommand InsertSessionCommand,
	selectSessionByAccessTokenCommand SelectSessionByAccessTokenCommand,
	updateSessionByAccessTokenCommand UpdateSessionByAccessTokenCommand,
	deleteSessionByAccessTokenCommand DeleteSessionByAccessTokenCommand) SessionRepository {
	return &sessionRepository{
		insertSessionCommand:              insertSessionCommand,
		selectSessionByAccessTokenCommand: selectSessionByAccessTokenCommand,
		updateSessionByAccessTokenCommand: updateSessionByAccessTokenCommand,
		deleteSessionByAccessTokenCommand: deleteSessionByAccessTokenCommand,
	}
}

func (s sessionRepository) Insert(ctx context.Context, session entities.Session) error {
	return s.insertSessionCommand.Execute(ctx, session)
}
