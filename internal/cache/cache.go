package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Cache struct{ client *redis.Client }

func New(addr string) *Cache  { return &Cache{redis.NewClient(&redis.Options{Addr: addr})} }
func (c *Cache) Close() error { return c.client.Close() }
func (c *Cache) Set(ctx context.Context, k, v string, ttl time.Duration) error {
	return c.client.Set(ctx, k, v, ttl).Err()
}
func (c *Cache) Get(ctx context.Context, k string) (string, error) {
	return c.client.Get(ctx, k).Result()
}
func (c *Cache) TTL(ctx context.Context, k string) (time.Duration, error) {
	return c.client.TTL(ctx, k).Result()
}
