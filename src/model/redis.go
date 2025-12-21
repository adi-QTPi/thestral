package model

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisClient(host, port, password string) (*RedisStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})

	store := &RedisStore{
		client: client,
	}
	return store, nil
}

func (rs *RedisStore) Add() {

}

func (rs *RedisStore) GetAll() {

}
