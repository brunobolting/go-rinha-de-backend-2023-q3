package person

import entity "github.com/brunobolting/go-rinha-backend/domain"

type Repository interface {
	Get(id string) (*entity.Person, error)
	Find(q string) ([]*entity.Person, error)
	Create(e *entity.Person) (string, error)
	Count() (int32, error)
}

type CacheRepository interface {
	Get(id string) (*entity.Person, error)
	Create(e *entity.Person) (string, error)
}

type UseCase interface {
	GetPerson(id string) (*entity.Person, error)
	FindPerson(q string) ([]*entity.Person, error)
	CreatePerson(nickname, name, birthdate string) (string, error)
	CountPerson() (int32, error)
}
