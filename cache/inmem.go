package cache

import (
	"errors"
	goc "github.com/pmylund/go-cache"
	"time"
)

// Stores data in memory. Major use-case is for use on development machines
type memStore struct {
	SimpleKeyFormat
	ram *goc.Cache
}

func NewInMemory() Cacher {
	return memStore{
		ram: goc.New(60*time.Minute, 1*time.Minute),
	}
}

func InMemoryFromConfig(container string) Cacher {
	return NewInMemory()
}

func (c memStore) Set(key string, data []byte, expireIn time.Duration) {
	key = c.Format(key)
	c.ram.Set(key, data, expireIn)
}

func (c memStore) Get(key string) ([]byte, error) {
	key = c.Format(key)
	if i, ok := c.ram.Get(key); ok {
		return i.([]byte), nil
	} else {
		return nil, errors.New("key not found in inmem cache: " + key)
	}
}

func (c memStore) Delete(key string) error {
	key = c.Format(key)
	c.ram.Delete(key)
	return nil
}

func (c memStore) Close() {

}
