package cache

import (
	"context"
	"time"
)

//go:generate mockgen -destination=mock/cache.go -package=cachemock -source=cache.go

// Cache defines commonly used cache interface
type Cache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, dest interface{}) error
	Del(ctx context.Context, keys ...string) error
}
