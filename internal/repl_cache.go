package internal

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	// Implement loop to clear cache
	c := Cache{
		Elements: make(map[string]CacheEntry),
		Mutex:    &sync.Mutex{},
	}

	go c.ReapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.Mutex.Lock() // We need the Locks cause maps aren't thread safe.
	ce := CacheEntry{
		CratedAt: time.Now().UTC(),
		Val:      val,
	}
	c.Elements[key] = ce
	defer c.Mutex.Unlock() // We unlock de lock to allow other threads to use the lock
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	ce, ok := c.Elements[key] // ce == CacheEntry
	return ce.Val, ok
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	for k, v := range c.Elements {
		if v.CratedAt.Before(now.Add(-last)) {
			delete(c.Elements, k)
		}
	}
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}
