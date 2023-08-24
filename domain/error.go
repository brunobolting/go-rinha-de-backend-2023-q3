package entity

import "errors"

var ErrInvalidNickname = errors.New("nickname is invalid")

var ErrInvalidName = errors.New("name is invalid")

var ErrEntityNotFound = errors.New("entity not found")
