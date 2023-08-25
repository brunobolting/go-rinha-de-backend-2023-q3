package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

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
	v, err := r.cache.Get(context.Background(), id).Result()

	if errors.Is(err, redis.Nil) {
		return nil, entity.ErrEntityNotFound
	}
	if err != nil {
		return nil, err
	}

	var p entity.Person

	err = json.Unmarshal([]byte(v), &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PersonRedis) Create(p *entity.Person) error {
	v, err := json.Marshal(p)
	if err != nil {
		return err
	}

	return r.cache.Set(context.Background(), p.ID.String(), v, 0).Err()
}

func (r *PersonRedis) SetNickname(n string) error {
	return r.cache.Set(context.Background(), fmt.Sprintf("nickname:%s", n), "", 0).Err()
}

func (r *PersonRedis) NicknameExists(n string) (bool, error) {
	v, err := r.cache.Exists(context.Background(), fmt.Sprintf("nickname:%s", n)).Result()
	return v == 1, err
}
