package cmaestro_db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	client *redis.Client
	ctx    context.Context
}

type Config struct {
	Addr     string
	Password string
	DB       int
}

func New(cfg Config) (*RedisDB, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx := context.Background()

	// Test connection
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisDB{
		client: rdb,
		ctx:    ctx,
	}, nil
}

// Close closes the Redis connection.
func (r *RedisDB) Close() error {
	return r.client.Close()
}

// Set stores a string value.
func (r *RedisDB) Set(key string, value string, expiration time.Duration) error {
	return r.client.Set(r.ctx, key, value, expiration).Err()
}

// Get retrieves a string value.
func (r *RedisDB) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

// Delete removes a key.
func (r *RedisDB) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

// Exists checks if a key exists.
func (r *RedisDB) Exists(key string) (bool, error) {
	count, err := r.client.Exists(r.ctx, key).Result()
	return count > 0, err
}

// SetJSON stores a struct/object as JSON.
func (r *RedisDB) SetJSON(key string, value any, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(r.ctx, key, data, expiration).Err()
}

// GetJSON retrieves JSON into a struct/object.
func (r *RedisDB) GetJSON(key string, dest any) error {
	data, err := r.client.Get(r.ctx, key).Bytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, dest)
}

// Increment increments an integer key.
func (r *RedisDB) Increment(key string) (int64, error) {
	return r.client.Incr(r.ctx, key).Result()
}

// Expire sets expiration for a key.
func (r *RedisDB) Expire(key string, expiration time.Duration) error {
	return r.client.Expire(r.ctx, key, expiration).Err()
}

// FlushDB clears the current database.
func (r *RedisDB) FlushDB() error {
	return r.client.FlushDB(r.ctx).Err()
}
