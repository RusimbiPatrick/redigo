package storage

import (
    "strconv"
    "sync"
    "time"
)

type Engine struct {
	mu sync.RWMutex
	items map[string]Item
	ttlHeap *TTLHeap
}

type Item struct {
    Key      string
    ExpireAt time.Time
    Value    interface{}
}

func NewEngine() *Engine {
	return &Engine{
		items: make(map[string]Item),
		ttlHeap: &TTLHeap{},
	}
}

func (e *Engine) Set(key, value string) []byte {
    e.mu.Lock()
    defer e.mu.Unlock()
    e.items[key] = Item{
        Key:      key, // Assign the key to the Item's Key field
        Value:    value,
        ExpireAt: time.Now().Add(time.Hour * 24 * 365), // Default 1 year
    }
    return []byte("+OK\r\n")
}

func (e *Engine) Get(key string) []byte {
    e.mu.RLock() // Acquire a read lock
    defer e.mu.RUnlock()

    item, exists := e.items[key]
    if !exists || time.Now().After(item.ExpireAt) {
        return []byte("$-1\r\n")
    }

    value, ok := item.Value.(string) // Ensure the value is a string
    if !ok {
        return []byte("-ERR value is not a string\r\n")
    }

    return []byte("$" + strconv.Itoa(len(value)) + "\r\n" + value + "\r\n")
}