package repository

import (
	"context"
	"encoding/json"
	"errors"

	entity "github.com/brunobolting/go-rinha-backend/domain"
	"github.com/go-redis/redis/v8"
)

type PersonRedis struct {
	cache *redis.Client
}

func NewPersonRedis(cache *redis.Client) *PersonRedis {
	return &PersonRedis{
		cache,
	}
}

func (r *PersonRedis) Get(id string) (*entity.Person, error) {
	v, err := r.cache.Get(context.Background(), id).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, entity.ErrEntityNotFound
	}
	if err != nil {
		return nil, err
	}

	var p *entity.Person

	err = json.Unmarshal(v, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *PersonRedis) Create(p *entity.Person) (string, error) {
	v, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	key := p.ID.String()
	err = r.cache.Set(context.Background(), key, v, 0).Err()
	if err != nil {
		return "", err
	}

	return key, nil
}
