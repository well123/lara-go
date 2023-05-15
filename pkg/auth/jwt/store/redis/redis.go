package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type Config struct {
	Addr     string
	DB       int
	Password string
	Prefix   string
}
type redisClient interface {
	Get(key string) *redis.StringCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Expire(key string, expire time.Duration) *redis.BoolCmd
	Exists(keys ...string) *redis.IntCmd
	Del(keys ...string) *redis.IntCmd
	TxPipeline() redis.Pipeliner
	Close() error
}
type Store struct {
	cli    redisClient
	prefix string
}

func NewStore(cfg *Config) *Store {
	cli := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &Store{
		cli:    cli,
		prefix: cfg.Prefix,
	}
}

func NewStoreWithClient(client *redis.Client, prefix string) *Store {
	return &Store{
		cli:    client,
		prefix: prefix,
	}
}

func NewStoreWithClusterClient(client *redis.ClusterClient, prefix string) *Store {
	return &Store{
		cli:    client,
		prefix: prefix,
	}
}

func (store *Store) wrapperKey(key string) string {
	return fmt.Sprintf("%s%s", store.prefix, key)
}

func (store *Store) Check(ctx context.Context, tokenString string) (bool, error) {
	cmd := store.cli.Exists(store.wrapperKey(tokenString))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

func (store *Store) Delete(ctx context.Context, tokenString string) (bool, error) {
	cmd := store.cli.Del(store.wrapperKey(tokenString))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

func (store *Store) Set(ctx context.Context, tokenString string, expiration time.Duration) error {
	cmd := store.cli.Set(store.wrapperKey(tokenString), "1", expiration)
	return cmd.Err()
}

func (store *Store) Release() error {
	return store.cli.Close()
}
