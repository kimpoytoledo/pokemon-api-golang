package errorhandler

import (
	"errors"
	"net/http"
	"pokemon-api-golang-v2/util/logger"
)

type ResponseError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

var (
	ErrPokemonNotFound = errors.New("Pokemon not found")
)

func HandleError(log *logger.Logger, err error) *ResponseError {
	log.Error(err)

	switch err {
	case ErrPokemonNotFound:
		return &ResponseError{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}
	default:
		return &ResponseError{
			Code:    http.StatusInternalServerError,
			Message: "An unexpected error occurred",
		}
	}
}
