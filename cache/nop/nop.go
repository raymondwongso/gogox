package nop

import (
	"context"
	"time"

	"github.com/raymondwongso/gogox/errorx"
)

// Nop implements Cache interface with no operation, commonly used for testing
type Nop struct{}

// New creates new Nop cache
func New() *Nop {
	return &Nop{}
}

// Get always returns errorx.ErrNotFound
func (n *Nop) Get(ctx context.Context, key string, dest interface{}) error {
	return errorx.ErrNotFound("no operation cache")
}

// Set do nothing and return nil error
func (n *Nop) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return nil
}

// Del do nothing and return nil error
func (n *Nop) Del(ctx context.Context, keys ...string) error {
	return nil
}
