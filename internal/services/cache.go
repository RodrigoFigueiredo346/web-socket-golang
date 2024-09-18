package services

import (
	"sync"
)

var cache *Cache
var once sync.Once

func init() {
	NewCache()
}

type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewCache() {
	once.Do(func() {
		cache = &Cache{
			data: make(map[string]interface{}),
		}
	})
}

func GetCache() *Cache {
	if cache == nil {
		NewCache()
	}
	return cache
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.data[key]
	return value, ok
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]interface{})
}

// // Adding data to the cache
// cache.Set("key1", "value1")
// cache.Set("key2", 123)

// // Retrieving data from the cache
// if val, ok := cache.Get("key1"); ok {
// 	fmt.Println("Value for key1:", val)
// }

// // Deleting data from the cache
// cache.Delete("key2")

// // Clearing the cache
// cache.Clear()

// time.Sleep(time.Second) // Sleep to allow cache operations to complete
