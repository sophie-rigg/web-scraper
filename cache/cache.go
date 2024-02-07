package cache

import (
	"sync"
)

type Client struct {
	sync.Map
}

func New() *Client {
	return &Client{}
}

// Set adds a key to the cache
func (c *Client) Set(key string) {
	c.Store(key, struct{}{})
}

// Exists returns true if the key exists in the cache
func (c *Client) Exists(key string) bool {
	_, ok := c.Load(key)
	return ok
}

// GetAll returns all keys in the cache
func (c *Client) GetAll() string {
	var keys string
	c.Range(func(key, value interface{}) bool {
		keys += key.(string) + "\n"
		return true
	})
	return keys
}
