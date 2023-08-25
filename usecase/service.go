package person

import (
	"errors"

	entity "github.com/brunobolting/go-rinha-backend/domain"
	"github.com/lib/pq"
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

func (s *Service) CreatePerson(nickname, name, birthdate string, stack []string) (*entity.Person, error) {
	e, err := entity.NewPerson(nickname, name, birthdate, stack)
	if err != nil {
		return nil, err
	}
	exists, err := s.cache.NicknameExists(e.Nickname)
	if exists {
		return nil, entity.ErrInvalidNickname
	}
	err = s.db.Create(e)
	if err != nil {
		if pgerr, ok := err.(*pq.Error); ok {
			if pgerr.Code == "23505" {
				return nil, entity.ErrInvalidNickname
			}
		}

		return nil, err
	}

	err = s.cache.Create(e)
	if err != nil {
		return nil, err
	}

	err = s.cache.SetNickname(e.Nickname)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (s *Service) CountPerson() (int32, error) {
	return s.db.Count()
}
