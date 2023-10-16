package usecase

import (
	"errors"
	"pokemon-api-golang-v2/core/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetByName(name string) (*entity.Pokemon, error) {
	args := m.Called(name)
	return args.Get(0).(*entity.Pokemon), args.Error(1)
}

func (m *MockRepository) Save(name string, pokemon *entity.Pokemon) error {
	args := m.Called(name, pokemon)
	return args.Error(0)
}

func TestFetchAndCachePokemon(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewPokemonService(mockRepo)

	pokemon := &entity.Pokemon{Name: "bulbasaur"}

	// Test case: Pokemon exists in cache
	mockRepo.On("GetByName", "bulbasaur").Return(pokemon, nil)
	result, err := usecase.FetchAndCachePokemon("bulbasaur")
	assert.NoError(t, err)
	assert.Equal(t, pokemon, result)

	// Test case: Pokemon not in cache
	mockRepo.On("GetByName", "bulbasaur").Return(nil, errors.New("not found"))
	mockRepo.On("Save", "bulbasaur", pokemon).Return(nil)
	result, err = usecase.FetchAndCachePokemon("bulbasaur")
	assert.NoError(t, err)
	assert.Equal(t, pokemon, result)
}
