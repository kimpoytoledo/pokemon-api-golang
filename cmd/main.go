package main

import (
	"pokemon-api-golang-v2/adapter/redis"
	"pokemon-api-golang-v2/api"
	"pokemon-api-golang-v2/core/usecase"

	"github.com/gin-gonic/gin"
	goredis "github.com/go-redis/redis/v8"
)

func main() {
	rdb := goredis.NewClient(&goredis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	repo := redis.NewRedisRepository(rdb)
	usecaseService := usecase.NewPokemonService(repo)
	handler := api.NewHandler(usecaseService)

	r := gin.Default()
	r.GET("/pokemon/:name", handler.GetPokemon)
	r.Run(":8080")
}
