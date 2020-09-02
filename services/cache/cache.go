package cache

import (
	"errors"
	"time"
)

// ErrKeyNotFound is the error when the given key is not found
var KeyNotFoundErr = errors.New("key not found in storage")

type Cache interface {
	// Set adds the given key and value to the store without an expiry. If the key already exists,
	// it will overwrite the previous value.
	Set(key string, value interface{}) error
	// SetWithExpiry adds the given key and value to the cache with the given expiry. If the key
	// already exists, it will overwrite the previoous value
	SetWithExpiry(key string, value interface{}, ttl time.Duration)
	// Get the content stored in the cache for the given key, and decode it into the value interface.
	// BUG(value): required pointer
	// Return error if the key is missing from the cache
	Get(key string, value interface{}) error
	// Remove deletes the key and value from storage
	Remove(key string) error

	// Purge all data and recreate cache service
	Purge() error
	// Len returns the number of items in the cache.
	Len() (int, error)
	// Keys returns a slice of the keys in the cache.
	Keys() ([]string, error)
}
