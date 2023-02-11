package memcache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

// Memcache implements Cache interface for memcache memory storage
type Memcache struct {
	client *memcache.Client
}

// New creates new Memcache implementation
func New(client *memcache.Client) *Memcache {
	return &Memcache{client: client}
}

// Get fetch key value and then unmarshal it into dest. Will return gogox error with code CodeNotFound if key not found.
func (m *Memcache) Get(ctx context.Context, key string, dest interface{}) error {
	it, err := m.client.Get(key)
	if err != nil {
		return err
	}

	return json.Unmarshal(it.Value, dest)
}

// Set set value for key using value json marshal result. Expiration is set accordingly
func (m *Memcache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	setBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return m.client.Set(&memcache.Item{Key: key, Value: setBytes, Expiration: int32(expiration.Seconds())})
}

// Del delete keys
func (m *Memcache) Del(ctx context.Context, keys ...string) error {
	for _, key := range keys {
		if err := m.client.Delete(key); err != nil {
			return err
		}
	}

	return nil
}
