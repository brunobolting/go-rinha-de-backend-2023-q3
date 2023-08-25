package entity

import "github.com/google/uuid"

type Person struct {
	ID        uuid.UUID `json:"id"`
	Nickname  string    `json:"apelido"`
	Name      string    `json:"nome"`
	Birthdate string    `json:"nascimento"`
	Stack     []string  `json:"stack"`
}

func NewPerson(nickname, name, birthdate string, stack []string) (*Person, error) {
	if !validateNickname(nickname) {
		return nil, ErrInvalidNickname
	}

	if !validateName(name) {
		return nil, ErrInvalidName
	}

	if !validateBirthdate(birthdate) {
		return nil, ErrInvalidBirthdate
	}

	if !validateStack(stack) {
		return nil, ErrInvalidStack
	}

	return &Person{
		ID:        uuid.New(),
		Nickname:  nickname,
		Name:      name,
		Birthdate: birthdate,
		Stack:     stack,
	}, nil
}

func validateNickname(nickname string) bool {
	return len(nickname) <= 32
}

func validateName(name string) bool {
	return len(name) <= 100
}

func validateBirthdate(birthdate string) bool {
	return len(birthdate) <= 10
}

func validateStack(stack []string) bool {
	for _, s := range stack {
		if len(s) > 32 {
			return false
		}
	}
	return true
}
