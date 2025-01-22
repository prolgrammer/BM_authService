package entities

type Account struct {
	Email    string
	Password string
}

func CreateAccount(email string, password string) *Account {
	return &Account{email, password}
}

func (a Account) Validate() error {
	return nil
}
