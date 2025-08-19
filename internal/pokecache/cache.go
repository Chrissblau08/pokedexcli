package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data     map[string]cacheEntry
	mutex    sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data:     make(map[string]cacheEntry),
		interval: interval,
	}

	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	//fmt.Println("[Cache] Added:", key)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	//fmt.Println("[Cache] Hit:", key)
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mutex.Lock()
		now := time.Now()
		for k, entry := range c.data {
			if now.Sub(entry.createdAt) > c.interval {
				delete(c.data, k)
				//fmt.Println("[Cache] Reaped:", k)
			}
		}
		c.mutex.Unlock()
	}
}
