package entity

import "errors"

var ErrInvalidNickname = errors.New("nickname is invalid")

var ErrInvalidName = errors.New("name is invalid")

var ErrEntityNotFound = errors.New("entity not found")

var ErrInvalidBirthdate = errors.New("birthdate is invalid")

var ErrInvalidStack = errors.New("stack is invalid")
