package usecases

import (
	"auth/controllers/requests"
	"auth/internal/entities"
	"context"
)

type signUpUseCase struct {
}

type SignUpUseCase interface {
	SignUp(ctx context.Context, request requests.SignRequest) (requests.SignRequest, error)
}

func NewSignUpUseCase() SignUpUseCase {
	return &signUpUseCase{}
}

func (u *signUpUseCase) SignUp(ctx context.Context, request requests.SignRequest) (requests.SignRequest, error) {
	account := entities.CreateAccount(request.Email, request.Password)
	err := account.Validate()
	if err != nil {
		return requests.SignRequest{}, err
	}

	return requests.SignRequest{}, nil
}
