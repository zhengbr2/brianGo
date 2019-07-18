
package main

import (
"fmt"
"runtime"
"time"
)



// Map provides an expiring storage for keys and values.
type Map struct {
	data       map[string]expiringValue
	expiration time.Duration
	done       chan struct{}
}

// expiringValue associates on piece of data with an expiration.
type expiringValue struct {
	expiration time.Time
	data       []byte
}

// NewMap creates a Map type and starts a worker goroutine to expire stale keys.
func NewMap(expiration time.Duration) *Map {
	m := &Map{
		data:       make(map[string]expiringValue),
		expiration: expiration,
		done:       make(chan struct{}),
	}

	go func() {
		ticker := time.NewTicker(expiration)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				m.removeExpired()
			case <-m.done:
				return
			}
		}
	}()

	return m
}

// Get retrieves a particular key from the map.
func (m *Map) Get(key string) ([]byte, bool) {
	v, ok := m.data[key]
	return v.data, ok
}

// Set records a key and value with the configuration expiration.
func (m *Map) Set(key string, value []byte) {
	expiration := time.Now().Add(m.expiration)
	m.data[key] = expiringValue{data: value, expiration: expiration}
}

// removeExpired removes any stale keys.
func (m *Map) removeExpired() {
	for k, v := range m.data {
		if time.Now().After(v.expiration) {
			delete(m.data, k)
		}
	}
}

// Close shoudl be called after the Map will no longer be used.
func (m *Map) Close() error {
	close(m.done)
	return nil
}

func main() {
	// This goroutine reports on the runtime's number of goroutines and the
	// memory allocator's heap allocations every five seconds.
	go func() {
		var stats runtime.MemStats
		for {
			runtime.ReadMemStats(&stats)
			fmt.Printf("HeapAlloc    = %d\n", stats.HeapAlloc)
			fmt.Printf("NumGoroutine = %d\n", runtime.NumGoroutine())
			time.Sleep(5 * time.Second)
		}
	}()

	// A tight loop to create a large number of ttl.Map types
	for {
		work()
	}
}

// work creates a map, sets a key, and then closes it down.
func work() {
	m := NewMap(5 * time.Minute)
	m.Set("my-key", []byte("my-value"))

	if _, ok := m.Get("my-key"); !ok {
		panic("no value present")
	}
	//m.Close()
	// m goes out of scope
}

