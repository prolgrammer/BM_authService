package usecases

import (
	"auth/controllers/requests"
	"auth/controllers/responses"
	"auth/internal/entities"
	"auth/internal/repositories"
	"context"
	"errors"
)

type signInUseCase struct {
	accountRepository repositories.AccountRepository
}

type SignInUseCase interface {
	SignIn(ctx context.Context, request requests.SignRequest) (responses.SignResponse, error)
}

func NewSignInUseCase(accountRepository repositories.AccountRepository) SignInUseCase {
	return &signInUseCase{accountRepository: accountRepository}
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

	return responses.SignResponse{
		Id: ac.Id,
	}, nil
}
