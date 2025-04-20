
package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(redisURL string) (*RedisCache, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &RedisCache{client: client}, nil
}

func (c *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(context.Background(), key, data, expiration).Err()
}

func (c *RedisCache) Get(key string, value interface{}) error {
	data, err := c.client.Get(context.Background(), key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, value)
}

func (c *RedisCache) Delete(key string) error {
	return c.client.Del(context.Background(), key).Err()
}

func (c *RedisCache) Clear() error {
	return c.client.FlushAll(context.Background()).Err()
}
