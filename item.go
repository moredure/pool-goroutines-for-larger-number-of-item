package main

import (
	"fmt"
	"time"
)

// all acounts share nothing
type Account struct {
	id             string
	dataArrived    bool
	stopped        chan struct{}
	state          State
	nextAccessTime time.Time
}

type State int

const (
	Connect State = iota
	Work
)

func (a *Account) Stop() {
	close(a.stopped)
}

func (a *Account) Connect() (State, error) {
	fmt.Println("Connect", a.id)
	return Work, nil
}

func (a *Account) Work() (State, error) {
	fmt.Println("Work", a.id)
	return Work, nil
}
