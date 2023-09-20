package redis

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"logur.dev/logur"
)

type Client struct {
	logger logur.LoggerFacade
	redis  *redis.Client
}

func New(ctx context.Context, address string, logger logur.LoggerFacade) (*Client, error) {
	r := redis.NewClient(&redis.Options{
		Addr: address,
	})
	if _, err := r.Ping(ctx).Result(); err != nil {
		return &Client{}, errors.WithStack(err)
	}
	logger.Info("Redis connected!")
	return &Client{
		redis:  r,
		logger: logger,
	}, nil
}

func TestClient(ctx context.Context, logger logur.LoggerFacade) *Client {
	c, err := New(ctx, "127.0.0.1:6379", logger)
	if err != nil {
		panic(err)
	}
	return c
}
