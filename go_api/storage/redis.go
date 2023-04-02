package storage

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"
	"strconv"
	"time"

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
	log.Println("redis GetUser check")
	redisCmd := r.client.Get(ctx, strconv.Itoa(userId))
	userBytes, err := redisCmd.Bytes()
	if err != nil {
		log.Println("redis GetUser err")
		return nil, err

	}

	b := bytes.NewReader(userBytes)
	var user m.User

	if err := gob.NewDecoder(b).Decode(&user); err != nil {
		log.Println("redis GetUser gob err")

		return nil, err
	}

	log.Println("printng user from Redis GetUser")
	log.Println(user)

	return &user, nil
}
func (r *RedisCache) SetUser(ctx context.Context, user *m.User) error {
	log.Println("redis SetUser hit")
	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(user); err != nil {
		return err
	}

	return r.client.Set(ctx, strconv.Itoa(user.Id), b.Bytes(), 25*time.Second).Err()
}
