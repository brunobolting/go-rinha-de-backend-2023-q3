package person

import (
	"errors"

	entity "github.com/brunobolting/go-rinha-backend/domain"
)

type Service struct {
	db    Repository
	cache CacheRepository
}

func NewService(db Repository, cache CacheRepository) *Service {
	return &Service{
		db,
		cache,
	}
}

func (s *Service) GetPerson(id string) (*entity.Person, error) {
	e, err := s.cache.Get(id)
	if errors.Is(err, entity.ErrEntityNotFound) {
		p, err := s.db.Get(id)
		if err != nil {
			return nil, err
		}
		s.cache.Create(p)
		return p, nil
	}
	return e, err
}

func (s *Service) FindPerson(q string) ([]*entity.Person, error) {
	return s.db.Find(q)
}

func (s *Service) CreatePerson(nickname, name, birthdate string) (string, error) {
	e, err := entity.NewPerson(nickname, name, birthdate)
	if err != nil {
		return "", err
	}
	v, err := s.db.Create(e)
	if err != nil {
		return "", err
	}

	s.cache.Create(e)

	return v, nil
}

func (s *Service) CountPerson() (int32, error) {
	return s.db.Count()
}
