package models

import (
	"errors"
)

var (
	ErrNoRecord           = errors.New("models: not matching record found")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)
