package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/adi-QTPi/thestral/src/types"
	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisClient(host, port, password string) (*RedisStorage, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %v", err)
	}
	log.Printf("Redis connection successful at %s:%s", host, port)

	store := &RedisStorage{
		client: client,
	}
	return store, nil
}

func (rs *RedisStorage) Add(proxyObj *types.ProxyRoute) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	data, err := json.Marshal(proxyObj)
	if err != nil {
		return fmt.Errorf("failed to marshal : %v", err)
	}

	err = rs.client.HSet(ctx, "routes", proxyObj.Source, data).Err()
	if err != nil {
		return fmt.Errorf("failed to insert into redis : %v", err)
	}

	return nil
}

func (rs *RedisStorage) GetAll() (map[string]*types.ProxyRoute, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	results, err := rs.client.HGetAll(ctx, "routes").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch from redis: %v", err)
	}

	routes := make(map[string]*types.ProxyRoute)

	for source, jsonVal := range results {
		var route types.ProxyRoute
		if err := json.Unmarshal([]byte(jsonVal), &route); err != nil {
			fmt.Printf("skipping corrupt route for %s: %v\n", source, err)
			continue
		}
		routes[source] = &route
	}

	return routes, nil
}
