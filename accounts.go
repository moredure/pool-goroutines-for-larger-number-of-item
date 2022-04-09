package main

import (
	"sync"
)

type Accounts struct {
	sync.Mutex
	accounts map[string]*Account
}

func (acs *Accounts) AddAccount(id string) {
	acs.Lock()
	defer acs.Unlock()
	a := &Account{
		id:          id,
		dataArrived: false, // for epoll should be atomic
		stopped:     make(chan struct{}),
	}
	acs.accounts[id] = a
	acheap.Push(a)
}

func (acs *Accounts) RemoveAccount(id string) {
	acs.Lock()
	defer acs.Unlock()
	acs.accounts[id].Stop()
}
