package main

import (
	"sync"
	"time"
)

type Accounts struct {
	sync.Mutex
	accounts map[string]*Account
}

func (acs *Accounts) AddAccount(id string) {
	acs.Lock()
	defer acs.Unlock()
	a := &Account{
		id:             id,
		dataArrived:    false, // epoll
		stopped:        make(chan struct{}),
		nextAccessTime: time.Time{},
	}
	acs.accounts[id] = a
	acheap.Push(a)
}

func (acs *Accounts) RemoveAccount(id string) {
	acs.Lock()
	defer acs.Unlock()

	acs.accounts[id].Stop()
}
