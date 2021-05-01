package cache

import (
	"errors"
	"github.com/google/wire"
	"time"
)

var (
	ProviderSet = wire.NewSet(NewRedisOptions, NewRedis)

	KeyNotFoundErr     = errors.New("key not found in storage")
	ValueMarshallErr   = errors.New("value marshall error")
	InvalidStorageErr  = errors.New("invalid storage")
	InternalStorageErr = errors.New("internal storage error")
)

type Cache interface {
	// Set adds the given key and value to the repository without an expiry. If the key already exists,
	// it will overwrite the previous value.
	Set(bucket string, key string, value interface{}) error
	// SetWithExpiry adds the given key and value to the cache with the given expiry. If the key
	// already exists, it will overwrite the previoous value
	SetWithExpiry(bucket string, key string, value interface{}, ttl time.Duration) error
	// Get the content stored in the cache for the given key, and decode it into the value interface.
	// BUG(value): required pointer
	// Return error if the key is missing from the cache
	Get(bucket string, key string, value interface{}) error
	// Remove deletes the key and value from storage
	Remove(bucket string, key string) error

	// Purge all data and recreate cache service
	Purge() error
	// Len returns the number of items in the cache.
	Len() (int, error)
	// Keys returns a slice of the keys in the cache.
	Keys() ([]string, error)
}
