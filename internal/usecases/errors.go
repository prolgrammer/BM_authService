package usecases

import "errors"

var (
	ErrEntityAlreadyExists = errors.New("entity already exists")
	ErrPasswordMismatch    = errors.New("password mismatch")
)
