package api

import (
	"net/http"
	"pokemon-api-golang-v2/core/port"
	"pokemon-api-golang-v2/util/errorhandler"
	"pokemon-api-golang-v2/util/logger"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Usecase port.PokemonUsecase
	log     *logger.Logger
}

func NewHandler(u port.PokemonUsecase, log *logger.Logger) *Handler {
	return &Handler{
		Usecase: u,
		log:     log,
	}
}

func (h *Handler) GetPokemon(c *gin.Context) {
	name := c.Param("name")
	pokemon, err := h.Usecase.FetchAndCachePokemon(name)

	if err != nil {
		responseError := errorhandler.HandleError(h.log, err)
		c.JSON(responseError.Code, gin.H{"error": responseError.Message})
		return
	}

	c.JSON(http.StatusOK, pokemon)
}
