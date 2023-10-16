package api

import (
	"net/http"
	"pokemon-api-golang-v2/core/port"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Usecase port.PokemonUsecase
}

func NewHandler(u port.PokemonUsecase) *Handler {
	return &Handler{Usecase: u}
}

func (h *Handler) GetPokemon(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pokemon name must be provided"})
		return
	}

	pokemon, err := h.Usecase.FetchAndCachePokemon(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if pokemon == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
		return
	}

	c.JSON(http.StatusOK, pokemon)
}
