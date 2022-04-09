package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// all acounts share nothing
type Account struct {
	//conn net.Conn
	id             string
	dataArrived    bool
	stopped        atomic.Value
	state          State
	nextAccessTime time.Time
}

type State int

const (
	Connect State = iota
	Work
)

func (a *Account) Stop() {
	a.stopped.Store(struct{}{})
}

func (a *Account) Connect() (State, error) {
	//conn, err := net.Dial("tcp", a.id)
	//if err != nil {
	//	return Connect, nil
	//}
	//epoll.Add(conn, func(conn) {
	//	a.dataArrived = true
	//}) ONESHOT
	fmt.Println("Connect", a.id)
	return Work, nil
}

func (a *Account) Work() (State, error) {
	//if a.dataArrived {
	//	// process arived data change state etc
	//  a.dataArrived = false
	// read all data
	// //epoll.Add(conn, func(conn) {
	//	//	a.dataArrived = true
	//	//}) ONESHOT
	//}
	fmt.Println("Work", a.id)
	return Work, nil
}
