package usecases

import (
	"auth/controllers/requests"
	"auth/controllers/responses"
	"auth/internal/entities"
	"auth/internal/repositories"
	"context"
	"errors"
	"fmt"
)

type signInUseCase struct {
	accountRepository repositories.AccountRepository
	sessionService    SessionService
	sessionRepository SessionRepository
}

type SignInUseCase interface {
	SignIn(ctx context.Context, request requests.SignRequest) (responses.SignResponse, error)
}

func NewSignInUseCase(accountRepository repositories.AccountRepository, sessionRepository repositories.SessionRepository, sessionService SessionService) SignInUseCase {
	return &signInUseCase{
		accountRepository: accountRepository,
		sessionService:    sessionService,
		sessionRepository: sessionRepository,
	}
}

func (s signInUseCase) SignIn(ctx context.Context, request requests.SignRequest) (responses.SignResponse, error) {
	ac, err := s.accountRepository.SelectByEmail(ctx, entities.Email(request.Email))
	if err != nil {
		if errors.Is(err, repositories.ErrEntityNotFound) {
			return responses.SignResponse{}, repositories.ErrEntityNotFound
		}
		return responses.SignResponse{}, err
	}

	if request.Password != ac.Password {
		return responses.SignResponse{}, ErrPasswordMismatch
	}

	session, err := s.sessionService.CreateSession(ac)
	if err != nil {
		return responses.SignResponse{}, fmt.Errorf("in create session error: %v", err)
	}

	err = s.sessionRepository.Insert(ctx, session)
	if err != nil {
		return responses.SignResponse{}, fmt.Errorf("in insert session error: %v", err)
	}

	return responses.SignResponse{
		Id:      ac.Id,
		Session: responses.NewSession(session.AccessToken, session.RefreshToken, session.ExpiresAt.Unix()),
	}, nil
}
