package redis

import (
	"encoding/json"
	"pokemon-api-golang-v2/core/entity"
	"testing"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func TestRedisRepository(t *testing.T) {
	db, mock := redismock.NewClientMock()
	repo := NewRedisRepository(db)

	pokemon := &entity.Pokemon{Name: "bulbasaur"}

	// Test Save method
	data, _ := json.Marshal(pokemon)
	mock.ExpectSet("bulbasaur", data, 0).SetVal("OK") // Added SetVal here
	err := repo.Save("bulbasaur", pokemon)
	assert.NoError(t, err)

	// Test GetByName method
	mock.ExpectGet("bulbasaur").SetVal(string(data))
	result, err := repo.GetByName("bulbasaur")
	assert.NoError(t, err)
	assert.Equal(t, pokemon.Name, result.Name)
}
