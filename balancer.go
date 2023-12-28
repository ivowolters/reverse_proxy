package main

import (
	"sync"
)

var endpoints = [3]string{
	"https://www.google.com/",
	"https://www.tweakers.net/",
	"https://www.nos.nl/",
}

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Rotate() int {
	c.mu.Lock()
	c.count += 1
	if c.count >= len(endpoints) {
		c.count = 0
	}
	count := c.count
	c.mu.Unlock()
	return count
}

var counter Counter

func fetchUrl() string {
	count := counter.Rotate()
	return endpoints[count]
}
