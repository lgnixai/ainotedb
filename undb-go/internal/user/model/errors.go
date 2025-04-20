package model

import "errors"

var (
	// ErrEmptyEmail is returned when email is empty
	ErrEmptyEmail = errors.New("email is required")
	// ErrEmptyPassword is returned when password is empty
	ErrEmptyPassword = errors.New("password is required")
	// ErrEmptyName is returned when name is empty
	ErrEmptyName = errors.New("name is required")
	// ErrUserNotFound is returned when user is not found
	ErrUserNotFound = errors.New("user not found")
	// ErrEmailAlreadyExists is returned when email already exists
	ErrEmailAlreadyExists = errors.New("email already exists")
	// ErrInvalidPassword is returned when password is invalid
	ErrInvalidPassword = errors.New("invalid password")
)
