package cache

import "github.com/redis/go-redis/v9"

func NewRedisClient(Addr string, Password string, DB int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
		DB:       DB,
	})
}
