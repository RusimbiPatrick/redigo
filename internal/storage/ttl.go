package storage

import (
	"container/heap"
	"time"
)

type TTLHeap []*Item

func (h TTLHeap) Len() int           { return len(h) }
func (h TTLHeap) Less(i, j int) bool { return h[i].ExpireAt.Before(h[j].ExpireAt) }
func (h TTLHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TTLHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *TTLHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (e *Engine) StartTTLChecker() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			e.mu.Lock()
			for e.ttlHeap.Len() > 0 {
				item := e.ttlHeap[0]
				if time.Now().Before(item.ExpireAt) {
					break
				}
				heap.Pop(e.ttlHeap)
				delete(e.items, item.Key)
			}
			e.mu.Unlock()
		}
	}()
}