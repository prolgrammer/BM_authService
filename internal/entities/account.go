package entities

import "time"

type Account struct {
	Id               string
	Email            string
	Password         string
	RegistrationDate time.Time
	Role             string
}

func NewAccount(email string, password string) Account {
	return Account{Email: email, Password: password, Role: UserRole, RegistrationDate: time.Now()}
}

func (a Account) Validate() error {
	return nil
}
