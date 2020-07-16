package ttlmap

import (
	"sync"
	"time"
)

type cacheItem struct {
	entryTime time.Time
	item      interface{}
}

// TTLMap type
type TTLMap struct {
	mu       sync.RWMutex
	duration time.Duration
	cache    map[interface{}]*cacheItem
}

// NewTTLMap comments
func NewTTLMap(duration time.Duration) *TTLMap {
	m := TTLMap{
		cache:    make(map[interface{}]*cacheItem),
		duration: duration,
	}

	go func() {
		for {
			time.Sleep(duration * 2)

			m.mu.Lock()

			for key, ci := range m.cache {
				if time.Since(ci.entryTime) > m.duration {
					delete(m.cache, key)
				}
			}

			m.mu.Unlock()

		}
	}()

	return &m
}

// Get comment
func (m *TTLMap) Get(key interface{}) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	ci, found := m.cache[key]
	if !found {
		return nil, false
	}

	if ci == nil {
		return nil, false
	}

	if time.Since(ci.entryTime) > m.duration {
		return nil, false
	}

	return ci.item, found
}

// Set sets a key value pair in my TTLMap
func (m *TTLMap) Set(key interface{}, value interface{}) {
	ci := &cacheItem{
		item:      value,
		entryTime: time.Now(),
	}
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache[key] = ci
}
