package redis

import (
	"context"
	"encoding/json"
	"pokemon-api-golang-v2/core/entity"
	"pokemon-api-golang-v2/core/port"

	"github.com/go-redis/redis/v8"
)

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) port.PokemonRepository {
	return &redisRepository{client: client}
}

func (r *redisRepository) GetByName(name string) (*entity.Pokemon, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, name).Result()
	if err != nil {
		return nil, err
	}
	var pokemon entity.Pokemon
	err = json.Unmarshal([]byte(val), &pokemon)
	return &pokemon, err
}

func (r *redisRepository) Save(name string, pokemon *entity.Pokemon) error {
	ctx := context.Background()
	data, err := json.Marshal(pokemon)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, name, data, 0).Err()
}
