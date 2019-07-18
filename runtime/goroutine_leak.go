package main

import "time"

func NewMap(expiration time.Duration) *Map {
	m := &Map{
		data:       make(map[string]expiringValue),
		expiration: expiration,
	}

	// start a worker goroutine
	go func() {
		for range time.Tick(expiration) {
			m.removeExpired()
		}
	}()

	return m
}

type expiringValue struct {
	expiration time.Time
	data       []byte //实际的值
}
