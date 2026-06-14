package core_cache

import (
	"crypto/sha256"
	"encoding/hex"
	"sync"
)

type Entry struct {
	data []byte
	etag string
}
type Cache struct {
	mu sync.RWMutex
	m  map[string]Entry
}

func NewCache() *Cache {
	return &Cache{
		m: map[string]Entry{},
	}
}
func (c *Cache) Get(key string) ([]byte, string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.m[key]
	if !exists {
		return nil, "", false
	}

	return append([]byte(nil), value.data...), value.etag, true
}

func (c *Cache) Set(key string, data []byte) string {
	etag := MakeETag(data)

	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[key] = Entry{
		data: append([]byte(nil), data...),
		etag: etag,
	}

	return etag
}

func MakeETag(data []byte) string {
	sum := sha256.Sum256(data)
	return `"` + hex.EncodeToString(sum[:16]) + `"`
}
