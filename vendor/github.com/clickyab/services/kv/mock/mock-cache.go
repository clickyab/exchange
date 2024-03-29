package mock

import (
	"bytes"
	"errors"
	"sync"
	"time"

	"github.com/clickyab/services/kv"
)

type cacheMock struct {
	data map[string][]byte
	lock sync.RWMutex
}

// Do is called to store the cache
func (ch *cacheMock) Do(e kv.Cacheable, t time.Duration) error {
	ch.lock.Lock()
	defer ch.lock.Unlock()
	name := e.String()
	target := &bytes.Buffer{}
	err := e.Decode(target)
	if err != nil {
		return err
	}

	ch.data[name] = target.Bytes()
	return nil
}

// Hit called when we need to load the cache
func (ch *cacheMock) Hit(key string, e kv.Cacheable) error {
	ch.lock.RLock()
	defer ch.lock.RUnlock()
	data, ok := ch.data[key]
	if !ok {
		return errors.New("not found")
	}

	buf := bytes.NewReader(data)
	return e.Encode(buf)
}

// GetData is used to get mock key for testing
func GetData(ch kv.CacheProvider, key string) (bool, []byte) {
	cm := ch.(*cacheMock)
	b, ok := cm.data[key]
	return ok, b
}

// NewCacheMock is the function to return the cache mock
func NewCacheMock() kv.CacheProvider {
	return &cacheMock{
		data: make(map[string][]byte),
		lock: sync.RWMutex{},
	}
}
