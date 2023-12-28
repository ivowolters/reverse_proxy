package main

import (
	"sync"
)

var endpoints = [2]string{
	"https://www.google.com/",
	"https://www.tweakers.net/",
}

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) rotate() int {
	c.mu.Lock()
	c.count += 1
	if c.count >= len(endpoints) {
		c.count = 0
	}
	c.mu.Unlock()
	return c.count
}

var counter Counter

func fetchUrl() string {
	count := counter.rotate()
	return endpoints[count]
}
