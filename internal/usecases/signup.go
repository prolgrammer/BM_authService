package usecases

import (
	"auth/controllers/requests"
	"auth/controllers/responses"
	"auth/internal/entities"
	"context"
	"fmt"
)

type signUpUseCase struct {
	accountRepository SignUpAccountRepository
}

type SignUpUseCase interface {
	SignUp(ctx context.Context, request requests.SignRequest) (responses.SignResponse, error)
}

func NewSignUpUseCase(
	accountRepository SignUpAccountRepository,
) SignUpUseCase {
	return &signUpUseCase{
		accountRepository: accountRepository,
	}
}

func (u *signUpUseCase) SignUp(ctx context.Context, request requests.SignRequest) (responses.SignResponse, error) {
	exists, err := u.accountRepository.CheckEmailExists(ctx, entities.Email(request.Email))
	if err != nil {
		return responses.SignResponse{}, fmt.Errorf("in check email error: %v", err)
	}
	if exists {
		return responses.SignResponse{}, ErrEntityAlreadyExists
	}

	account := entities.NewAccount(request.Email, request.Password)
	err = account.Validate()
	if err != nil {
		return responses.SignResponse{}, err
	}

	account.Id, err = u.accountRepository.Insert(ctx, account)
	if err != nil {
		return responses.SignResponse{}, fmt.Errorf("in insert account error: %v", err)
	}

	return responses.SignResponse{
		Id: account.Id,
	}, nil
}
