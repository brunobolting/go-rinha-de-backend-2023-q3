package entity

import "github.com/google/uuid"

type Person struct {
	ID        uuid.UUID `json:"id"`
	Nickname  string    `json:"apelido"`
	Name      string    `json:"nome"`
	Birthdate string    `json:"nascimento"`
	Stack     []string  `json:"stack"`
}

func NewPerson(nickname, name, birthdate string) (*Person, error) {
	if !validateNickname(nickname) {
		return nil, ErrInvalidNickname
	}

	if !validateName(name) {
		return nil, ErrInvalidName
	}

	return &Person{
		ID:        uuid.New(),
		Nickname:  nickname,
		Name:      name,
		Birthdate: birthdate,
	}, nil
}

func validateNickname(nickname string) bool {
	return len(nickname) > 32
}

func validateName(name string) bool {
	return len(name) > 100
}
