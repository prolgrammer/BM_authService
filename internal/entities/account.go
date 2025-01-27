package entities

import "time"

type Account struct {
	Id               string
	Email            Email
	Password         Password
	RegistrationDate time.Time
	isVerified       bool
	Role             string
}

func NewAccount(email string, password string) Account {
	return Account{Email: Email(email), Password: Password(password), Role: UserRole, RegistrationDate: time.Now()}
}

func (a Account) Validate() error {
	return nil
}
