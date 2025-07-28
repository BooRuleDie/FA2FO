package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"go-backend/internal/store"
	"time"

	"github.com/redis/go-redis/v9"
)

type usersCacheRepo interface {
	Get(ctx context.Context, userID int64) (*store.User, error)
	Set(ctx context.Context, user *store.User) error
}

type usersStore struct {
	rc *redis.Client
}

func (us *usersStore) getUserCacheKey(userID int64) string {
	return fmt.Sprintf("user-%v", userID)
}

func (us *usersStore) Get(ctx context.Context, userID int64) (*store.User, error) {
	cacheKey := us.getUserCacheKey(userID)

	data, err := us.rc.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var user store.User
	if data != "" {
		if err := json.Unmarshal([]byte(data), &user); err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (us *usersStore) Set(ctx context.Context, user *store.User) error {
	cacheKey := us.getUserCacheKey(user.ID)
	expiryTime := time.Minute

	jsonData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return us.rc.SetEx(ctx, cacheKey, jsonData, expiryTime).Err()
}
