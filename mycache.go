/**
 * 一个基于内存的缓存模块
 **/

package mycache

import (
	"sync"
	"time"
)

type Cache struct {
	lru *LRU
	mu  sync.RWMutex
}

func New(size int, duration time.Duration) (*Cache, error) {
	lru, err := NewLRU(size, duration)
	if err != nil {
		return nil, err
	}
	c := &Cache{
		lru: lru,
	}
	return c, nil
}

func (c *Cache) Get(k interface{}) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.lru.Get(k)
}

func (c *Cache) Set(k interface{}, v interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.lru.Set(k, v)
}

/**
 * 添加节点
 * 如果节点已经存在且未超时，返回false
 **/
func (c *Cache) SetEx(k interface{}, v interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.lru.SetEx(k, v)
}
