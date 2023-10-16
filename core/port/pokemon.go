package port

import "pokemon-api-golang-v2/core/entity"

type PokemonRepository interface {
	GetByName(name string) (*entity.Pokemon, error)
	Save(name string, pokemon *entity.Pokemon) error
}

type PokemonUsecase interface {
	FetchAndCachePokemon(name string) (*entity.Pokemon, error)
}
