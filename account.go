package main

import (
	"fmt"
	"time"
)

// all acounts share nothing
type Account struct {
	//conn net.Conn
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
	//conn, err := net.Dial("tcp", a.id)
	//if err != nil {
	//	return Connect, nil
	//}
	//epoll.Add(conn, func(conn) {
	//	a.dataArrived = true
	//})
	fmt.Println("Connect", a.id)
	return Work, nil
}

func (a *Account) Work() (State, error) {
	//if a.dataArrived {
	//	// process arived data change state etc
	//}
	fmt.Println("Work", a.id)
	return Work, nil
}
