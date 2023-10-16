package pokeapi

import (
	"encoding/json"
	"net/http"
	"pokemon-api-golang-v2/core/entity"
)

func FetchPokemonFromAPI(name string) (*entity.Pokemon, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Note: This is a simple extraction. The PokeAPI has a lot more data than just the name.
	var result struct {
		Name string `json:"name"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &entity.Pokemon{Name: result.Name}, nil
}
