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
	sessionRepository SessionRepository
	sessionService    SessionService
	hashService       SignUpHashService
}

type SignUpUseCase interface {
	SignUp(ctx context.Context, request requests.SignRequest) (responses.SignResponse, error)
}

func NewSignUpUseCase(
	accountRepository SignUpAccountRepository,
	sessionRepository SessionRepository,
	sessionService SessionService,
	hashService SignUpHashService,
) SignUpUseCase {
	return &signUpUseCase{
		accountRepository: accountRepository,
		sessionRepository: sessionRepository,
		sessionService:    sessionService,
		hashService:       hashService,
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

	hashedPassword, err := u.hashService.CreateHash(request.Password)
	if err != nil {
		return responses.SignResponse{}, fmt.Errorf("failed to hash the password: %v", err)
	}
	account.Password = entities.Password(hashedPassword)

	account.Id, err = u.accountRepository.Insert(ctx, account)
	if err != nil {
		return responses.SignResponse{}, fmt.Errorf("in insert account error: %v", err)
	}

	session, err := u.sessionService.CreateSession(account)
	if err != nil {
		return responses.SignResponse{}, fmt.Errorf("in create session error: %v", err)
	}

	err = u.sessionRepository.Insert(ctx, session)
	if err != nil {
		return responses.SignResponse{}, fmt.Errorf("in insert session error: %v", err)
	}

	return responses.SignResponse{
		Id:      account.Id,
		Session: responses.NewSession(session.AccessToken, session.RefreshToken, session.ExpiresAt.Unix()),
	}, nil
}
