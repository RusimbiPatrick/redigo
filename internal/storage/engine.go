package storage

import (
	"sync"
	"time"
)

type Engine struct {
	mu sync.RWMutex
	items map[string]Item
	ttlHeap *TTLHeap
}

type Item struct {
	Value string
	ExpireAt time.Time
}

func NewEngine() *Engine {
	return &Engine{
		items: make(map[string]Item),
		ttlHeap: &TTLHeap{},
	}
}

func(e *Engine) Set(Key, value string) []byte {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.items[key] = Item {
		Value: value,
		ExpireAt: time.Now().Add(time.Hour * 24 * 365), //Default 1 year
	}
	return []byte("+OK\r\n")
}

func(e *Engine) Get(key string) []byte {
	e.mu.Unlock()
	defer e.mu.RUnlock()

	item, exists := e.items[key];
	if !exists || time.Now().After(item.ExpireAt){
		return []byte("$-1\r\n")
	}
	return []byte("$" + string(len(item.Value)) + "\r\n" + item.Value + "\r\n")
}

