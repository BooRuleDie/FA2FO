package cache

import (
	"github.com/redis/go-redis/v9"
)

type CacheStorage struct {
	Users usersCacheRepo
}

func NewCacheStorage(rc *redis.Client) CacheStorage {
	return CacheStorage{
		Users: &usersStore{rc: rc},
	}
}
