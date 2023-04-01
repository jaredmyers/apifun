package storage

import (
	"context"
	"log"
	"strconv"

	m "github.com/jaredmyers/apifun/go_api/models"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache() (*RedisCache, error) {
	// to be implemented

	var ctx = context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	rc := &RedisCache{
		client: client,
	}
	return rc, nil
}

func (r *RedisCache) GetUser(ctx context.Context, userId int) (*m.User, error) {
	log.Println("GetUser from redis working")
	u, err := r.client.Get(ctx, strconv.Itoa(userId)).Result()
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *RedisCache) SetUser(ctx context.Context) error {
	log.Println("SetUser from redi working")
	return nil
}
