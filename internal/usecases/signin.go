package usecases

import (
	"context"
	"errors"
	"fmt"
	"github.com/prolgrammer/BM_authService/controllers/requests"
	"github.com/prolgrammer/BM_authService/controllers/responses"
	"github.com/prolgrammer/BM_authService/internal/repositories"
	e "github.com/prolgrammer/BM_package/errors"
)

type signInUseCase struct {
	accountRepository SignInAccountRepository
	sessionService    SessionService
	sessionRepository SessionRepository
	hashService       SignInHashService
}

type SignInUseCase interface {
	SignIn(ctx context.Context, request requests.SignRequest) (responses.SignResponse, error)
}

func NewSignInUseCase(accountRepository SignInAccountRepository, sessionRepository repositories.SessionRepository, sessionService SessionService, hashService SignInHashService) SignInUseCase {
	return &signInUseCase{
		accountRepository: accountRepository,
		sessionService:    sessionService,
		sessionRepository: sessionRepository,
		hashService:       hashService,
	}
}

func (s signInUseCase) SignIn(ctx context.Context, request requests.SignRequest) (responses.SignResponse, error) {
	ac, err := s.accountRepository.SelectByEmail(ctx, request.Email)
	if err != nil {
		if errors.Is(err, e.ErrEntityNotFound) {
			return responses.SignResponse{}, e.ErrEntityNotFound
		}
		return responses.SignResponse{}, err
	}

	_, err = s.hashService.CompareStringAndHash(request.Password, string(ac.Password))
	if err != nil {
		return responses.SignResponse{}, fmt.Errorf("failed to compare password: %w", err)
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
