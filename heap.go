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

type Accounts map[string]*Account

func (accounts Accounts) AddAccount(id string) {
	// grab from db
	a := &Account{
		stopped:        make(chan struct{}), // atomic
		dataArrived:    false,               // for epoll? to check non blocking whether the date arrived
		state:          0,
		nextAccessTime: time.Now(),
	}
	accounts[id] = a
	acheap.Push(a)
}

func (accounts Accounts) onStopAccount(id string) {
	accounts[id].Stop()
}
