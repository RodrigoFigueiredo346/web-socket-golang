package services

import (
	"fmt"
	"sync"
	"unsafe"
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

func (c *Cache) GetAll() map[string]interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Retorna uma cópia do mapa para evitar concorrência direta ao acessar o cache original
	copiedData := make(map[string]interface{})
	for key, value := range c.data {
		copiedData[key] = value
	}
	return copiedData
}

func (c *Cache) GetMemoryUsage() string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var totalSize uintptr
	for key, value := range c.data {
		totalSize += unsafe.Sizeof(key) + unsafe.Sizeof(value)
	}

	return fmt.Sprintf("Approximate memory usage: %d bytes", totalSize)
}
