package cache

import (
	"encoding/json"
	"fmt"
	redis_scan "github.com/gomodule/redigo/redis"
	"github.com/rl-os/api/services/redis"
	"time"
)

type RedisOptions struct{}

func NewRedisOptions() *RedisOptions {
	return &RedisOptions{}
}

type RedisCache struct {
	pool *redis.Pool

	options *RedisOptions
}

func NewRedis(options *RedisOptions, pool *redis.Pool) (Cache, error) {
	return &RedisCache{pool, options}, nil
}

func (c *RedisCache) Set(bucket string, key string, value interface{}) error {
	return c.SetWithExpiry(bucket, key, value, time.Minute)
}

func (c *RedisCache) SetWithExpiry(bucket string, key string, value interface{}, ttl time.Duration) error {
	keyName := fmt.Sprintf("app::cache::%s::%s", bucket, key)

	raw, err := json.Marshal(value)
	if err != nil {
		return ValueMarshallErr
	}

	_, err = c.pool.Get().Do("SET", keyName, raw, "EX", int(ttl.Seconds()))
	return err
}

func (c *RedisCache) Get(bucket string, key string, value interface{}) error {
	keyName := fmt.Sprintf("app::cache::%s::%s", bucket, key)

	raw, err := redis_scan.Bytes(c.pool.Get().Do("GET", keyName))
	if err != nil {
		return KeyNotFoundErr
	}

	err = json.Unmarshal(raw, value)
	if err != nil {
		return KeyNotFoundErr
	}

	return nil
}

func (c *RedisCache) Remove(bucket string, key string) error {
	panic("implement me")
}

func (c *RedisCache) Purge() error {
	panic("implement me")
}

func (c *RedisCache) Len() (int, error) {
	panic("implement me")
}

func (c *RedisCache) Keys() ([]string, error) {
	panic("implement me")
}
