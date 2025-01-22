package controllers

import "errors"

var (
	DataBindError      = errors.New("wrong data format")
	AuthRequiredError  = errors.New("auth is required")
	RegistrationsError = errors.New("registrations are required")
)
