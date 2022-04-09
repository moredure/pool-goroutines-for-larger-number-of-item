package main

import (
	"container/heap"
	"sync"
)

type SyncMutexItemsHeap struct {
	*sync.Cond
	Items Items
}

func (s *SyncMutexItemsHeap) Push(i *Account) {
	s.L.Lock()
	defer s.L.Unlock()
	heap.Push(&s.Items, i)
	s.Signal()
}

func (s *SyncMutexItemsHeap) Pop() *Account {
	s.L.Lock()
	defer s.L.Unlock()

	for len(s.Items) == 0 {
		s.Wait()
	}
	return heap.Pop(&s.Items).(*Account)
}
