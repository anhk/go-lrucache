/**
 *
 **/

package mycache

import (
	"container/list"
	"errors"
	"time"
)

/**
 * 缓存节点
 **/
type entry struct {
	key        interface{}
	value      interface{}
	expiration time.Time /** 到期时间戳, int64 **/
	lastAccess time.Time /** 最后访问时间 **/
}

/**
 * LRU结构定义
 **/
type LRU struct {
	size      int
	evictList *MyList
	items     map[interface{}]*list.Element
	duration  time.Duration /** 到期间隔，单位：秒 **/
}

func NewLRU(size int, duration time.Duration) (*LRU, error) {
	if size <= 0 {
		return nil, errors.New("Must provide a positive size")
	}
	c := &LRU{
		size:      size,
		evictList: &MyList{},
		items:     make(map[interface{}]*list.Element),
		duration:  duration * time.Second, /** 存活时间 **/
	}
	return c, nil
}

/**
 * 添加节点
 **/
func (c *LRU) Set(key, value interface{}) {
	var e time.Time
	if c.duration > 0 {
		e = time.Now().Add(c.duration)
	}

	/** 检查是否已经存在节点，如果存在则覆盖并置为最新 **/
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		ent.Value.(*entry).value = value
		ent.Value.(*entry).expiration = e
		return
	}

	/** 添加新的节点 **/
	ent := &entry{
		key:        key,
		value:      value,
		expiration: e,
		lastAccess: time.Now(),
	}
	entry := c.evictList.PushFront(ent)
	c.items[key] = entry

	/** 检查容量是否超限 **/
	if c.evictList.Len() > c.size {
		c.removeOldest()
	}
}

/**
 * 添加节点
 * 如果节点已经存在且未超时，返回false
 **/
func (c *LRU) SetEx(key, value interface{}) bool {
	if ent, ok := c.items[key]; ok {
		e := ent.Value.(*entry).expiration
		if e.IsZero() || time.Now().Before(e) { /** 没有过期 **/
			return false
		}
	}
	c.Set(key, value)
	return true
}

/**
 * 获取节点
 **/
func (c *LRU) Get(key interface{}) (value interface{}, err error) {
	if ent, ok := c.items[key]; ok {
		e := ent.Value.(*entry).expiration
		last := ent.Value.(*entry).lastAccess
		ent.Value.(*entry).lastAccess = time.Now()
		if !e.IsZero() && time.Now().After(e) { /** 已经过期 **/
			ent.Value.(*entry).expiration = time.Now().Add(c.duration)
			c.evictList.MoveToFront(ent)
			return ent.Value.(*entry).value, NewErrorExpired(last, "Expired.")
		}
		return ent.Value.(*entry).value, nil
	}
	return nil, NewErrorNonExisted("non-existed.")
}

/**
 * 删除最久未访问节点
 **/
func (c *LRU) removeOldest() {
	ent := c.evictList.RemoveBack()
	if ent != nil {
		kv := ent.(*entry)
		delete(c.items, kv.key)
	}
}
