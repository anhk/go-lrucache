package lrucache

import (
	"container/list"
	"sync"
)

type MyList struct {
	list list.List
	mu   sync.Mutex
}

func (l *MyList) MoveToFront(e *list.Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.MoveToFront(e)
}

func (l *MyList) PushFront(v interface{}) *list.Element {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.PushFront(v)
}

func (l *MyList) Remove(e *list.Element) interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.Remove(e)
}

func (l *MyList) RemoveBack() interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()
	ent := l.list.Back()
	if ent != nil {
		return l.list.Remove(ent)
	}
	return nil
}

func (l *MyList) Len() int {
	return l.list.Len()
}
