package repository

import (
	"context"
	"eth_bsc_multichain/internal/service/db/redis"
	"time"
)

type CacheRepo interface {
	Ping(ctx context.Context) error
	Get(ctx context.Context, key string, value interface{}) error
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
	HGet(ctx context.Context, hashKey string, values interface{}, emptyJson string, keys ...string) error
	HSet(ctx context.Context, hashKey string, data interface{}) error
}

func NewCacheRepoImpl(client *redis.Client) CacheRepo {
	return client
}
