package cli

import "errors"

var (
	ErrInvalidCommand   = errors.New("invaid command")
	ErrInvalidParameter = errors.New("invalid parameter")
)
