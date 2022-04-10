package main

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var acheap = &SyncMutexItemsHeap{
	Cond: sync.NewCond(new(sync.Mutex)),
}

var accounts = &Accounts{
	accounts: make(map[string]*Account),
}

var N = 2

func main() {
	go func() {
		accounts.AddAccount(strconv.Itoa(rand.Int()))
	}()
	var wg sync.WaitGroup

	// listen for nats on new accounts
	// listen for nats on remove account
	for i := 0; i < N; i += 1 { // N is number of workers
		wg.Add(1)
		// can be dynamicaly resized if size of acheap grows?
		go func() {
			defer wg.Done()
			for {
				account := acheap.Pop() // channel based pop can be used
				if pause := account.nextAccessTime.Sub(time.Now()); pause > 0 {
					time.Sleep(pause) // then it will be possible to wait on timer here and probably 
					// requeue sleeping item while starting newly appeared non sleeping item
				}
				if account.stopped.Load() != nil {
					continue
				}
				var (
					next State
					err  error
				)
				switch account.state {
				case Connect:
					// in case of error retry
					// in case of result return next State
					next, err = account.Connect()
				case Work:
					next, err = account.Work()
				default:
					log.Fatal("undefined State")
				}
				if err != nil {
					account.state = account.state
				} else {
					account.state = next
				}
				
				// optional
				account.nextAccessTime = time.Now().Add(2 * time.Second) // this time can be dynamicaly delayed
				acheap.Push(account)
				// also time.AfterFunc() can be done in order to prevent time.Sleep at the top blocking entire goroutine
				// other delayed enqueing can be used: some service with goroutine which checking all delayed tasks when they become ready
				// otherwise accounts can be pushed to acheap from epoll notified events to prevent unneeded processing of state
			}
		}()
	}
	wg.Wait()
}
