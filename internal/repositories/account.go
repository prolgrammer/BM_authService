package repositories

import (
	"context"
	"errors"
	"github.com/prolgrammer/BM_authService/internal/entities"
)

type accountRepository struct {
	insertAccountCommand        InsertAccountCommand
	selectAccountByEmailCommand SelectAccountByEmailCommand
}

type AccountRepository interface {
	Insert(ctx context.Context, account entities.Account) (string, error)
	//SelectById(ctx context.Context, id string) (entities.Account, error)
	SelectByEmail(ctx context.Context, email string) (entities.Account, error)
	CheckEmailExists(ctx context.Context, email entities.Email) (bool, error)
	//ChangePassword(ctx context.Context, id string, newPassword string) error
}

func NewAccountRepository(
	insertAccountCommand InsertAccountCommand,
	selectAccountByEmailCommand SelectAccountByEmailCommand,
) AccountRepository {
	return &accountRepository{
		insertAccountCommand:        insertAccountCommand,
		selectAccountByEmailCommand: selectAccountByEmailCommand,
	}
}

func (u *accountRepository) Insert(ctx context.Context, account entities.Account) (string, error) {
	return u.insertAccountCommand.Execute(ctx, account)
}

func (u *accountRepository) SelectByEmail(ctx context.Context, email string) (entities.Account, error) {
	return u.selectAccountByEmailCommand.Execute(ctx, entities.Email(email))
}

func (u *accountRepository) CheckEmailExists(ctx context.Context, email entities.Email) (bool, error) {
	_, err := u.selectAccountByEmailCommand.Execute(ctx, email)
	if err != nil {
		if errors.Is(err, ErrEntityNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
