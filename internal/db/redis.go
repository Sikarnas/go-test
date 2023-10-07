package db

import (
	"github.com/redis/go-redis/v9"
)

func New() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}
