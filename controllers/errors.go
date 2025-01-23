package controllers

import "errors"

var (
	ErrDataBindError      = errors.New("wrong data format")
	ErrAuthRequiredError  = errors.New("auth is required")
	ErrRegistrationsError = errors.New("registrations are required")
)
