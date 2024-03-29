package kv

import (
	"time"

	"github.com/clickyab/services/assert"
)

// AKiwi is the key value storage system in a parent key
type AKiwi interface {
	// Key return the parent key
	Key() string
	// IncSetSubKey for increasing sub key
	IncSubKey(key string, value int64) int64
	// DecSetSubKey for decreasing sub key
	DecSubKey(key string, value int64) int64
	// SubKey return a key
	SubKey(key string) int64
	// AllKeys from the store
	AllKeys() map[string]int64
	// Drop the entire eav store
	Drop() error
	// TTL return the time to expire this
	TTL() time.Duration
}

// StoreAtomicFactory is a function to create store
type StoreAtomicFactory func(string, time.Duration) AKiwi

var (
	atomicFactory StoreAtomicFactory
)

// NewAEAVStore return a new eav store
func NewAEAVStore(key string, duration time.Duration) AKiwi {
	regLock.RLock()
	defer regLock.RUnlock()

	assert.NotNil(atomicFactory, "[BUG] factory is not registered")
	return atomicFactory(key, duration)
}
