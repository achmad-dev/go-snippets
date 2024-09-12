package utils

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]interface{}
	mutex sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.cache[key]
}

// ttl is optional
func (c *Cache) Set(key string, value interface{}, ttl ...time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = value
	if len(ttl) > 0 {
		c.setTTL(key, ttl[0])
	}
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.cache, key)
}

func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache = make(map[string]interface{})
}

func (c *Cache) Exists(key string) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	_, ok := c.cache[key]
	return ok
}

// auto delete when TTL is expired
func (c *Cache) setTTL(key string, ttl time.Duration) {
	go func() {
		for {
			time.Sleep(ttl)
			c.Delete(key)
		}
	}()
}
