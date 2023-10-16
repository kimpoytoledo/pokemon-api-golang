package usecase

import (
	"pokemon-api-golang-v2/adapter/pokeapi"
	"pokemon-api-golang-v2/core/entity"
	"pokemon-api-golang-v2/core/port"
)

type pokemonService struct {
	repo port.PokemonRepository
}

func NewPokemonService(r port.PokemonRepository) port.PokemonUsecase {
	return &pokemonService{repo: r}
}

func (s *pokemonService) FetchAndCachePokemon(name string) (*entity.Pokemon, error) {
	pokemon, err := s.repo.GetByName(name)

	// If found in cache, return it
	if err == nil {
		pokemon.Source = "Redis Cache"
		return pokemon, nil
	}

	// If not in cache, fetch from PokeAPI
	pokemon, err = pokeapi.FetchPokemonFromAPI(name)
	if err != nil {
		return nil, err
	}
	pokemon.Source = "PokeAPI"

	// Save to cache
	err = s.repo.Save(name, &entity.Pokemon{Name: pokemon.Name})
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}
