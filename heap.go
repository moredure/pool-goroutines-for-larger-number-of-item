package main

import (
	"container/heap"
	"sync"
	"time"
)

type SyncMutexItemsHeap struct {
	sync.Mutex
	sync.Cond
	Items Items
}

func (s *SyncMutexItemsHeap) Push(i *Account) {
	s.Lock()
	defer s.Unlock()
	heap.Push(&s.Items, i)
	s.Signal()
}

func (s *SyncMutexItemsHeap) Pop() *Account {
	s.Lock()
	defer s.Unlock()

	for len(s.Items) == 0 {
		s.Wait()
	}
	return heap.Pop(&s.Items).(*Account)
}
