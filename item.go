package main

import (
	"time"
)

// all acounts share nothing
type Account struct {
	dataArrived    bool
	stopped        chan struct{}
	state          int
	nextAccessTime time.Time
}

func (a *Account) Stop() {
	close(a.stopped)
}

func (a *Account) SomeAction() (int, error) {
	// based on previous state and account
	return 0, nil
}

func (a *Account) SomeAction2() (int, error) {
	return 0, nil
}
