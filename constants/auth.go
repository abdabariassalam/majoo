package constants

import "errors"

var (
	ErrInvalidToken  = errors.New("token is invalid")
	ErrExpiredToken  = errors.New("token has expired")
	ErrTokenRequired = "token is required"
	ErrBadToken      = "bad token"
)
