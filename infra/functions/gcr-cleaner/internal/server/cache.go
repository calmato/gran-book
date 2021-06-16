package server

import (
	"sync"
	"time"
)

// Cache - is an interface used by the PubSub() function to prevent duplicate messages from being processed.
type Cache interface {
	// Insert - inserts the item into the cache.
	// If the item already exists, this method returns true.
	Insert(string) bool

	// Stop - stops  the cache.
	// When Stop returns, the cache must not perform any additionally proccessing.
	Stop()
}

// timerCache - is a Cache implementation that caches items for a configurable preod of time.
type timerCache struct {
	lock     sync.RWMutex
	data     map[string]struct{}
	lifetime time.Duration

	stopCh  chan struct{}
	stopped bool
}

// NewTimerCache - creates a new timer-based cache.
func NewTimerCache(lifetime time.Duration) Cache {
	return &timerCache{
		data:     make(map[string]struct{}),
		lifetime: lifetime,
		stopCh:   make(chan struct{}),
	}
}

// Insert - adds the item to the cache.
// If the item already existed in the cache, this function returns false.
func (c *timerCache) Insert(s string) bool {
	// Read only
	c.lock.RLock()
	if _, ok := c.data[s]; ok {
		c.lock.RUnlock()
		return true
	}
	c.lock.RUnlock()

	// Full insert
	c.lock.Lock()
	if _, ok := c.data[s]; ok {
		c.lock.Unlock()
		return true
	}

	c.data[s] = struct{}{}
	c.lock.Unlock()

	// Start a timeout to delete the item from the cache.
	go c.timeout(s)

	return false
}

// Stop - stops the cache.
func (c *timerCache) Stop() {
	c.lock.Lock()
	if !c.stopped {
		close(c.stopCh)
		c.stopped = true
	}
	c.lock.Unlock()
}

func (c *timerCache) timeout(s string) {
	select {
	case <-time.After(c.lifetime):
		c.lock.Lock()
		delete(c.data, s)
		c.lock.Unlock()
	case <-c.stopCh:
	}
}
