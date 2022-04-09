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
				account := acheap.Pop()
				if pause := account.nextAccessTime.Sub(time.Now()); pause > 0 {
					time.Sleep(pause)
				}
				select {
				case <-account.stopped:
					continue
				default:
					var err error
					var newState State
					switch account.state {
					case Connect:
						// in case of error retry
						// in case of result return next State
						newState, err = account.Connect()
					case Work:
						newState, err = account.Work()
					default:
						log.Fatal("undefined State")
					}
					if err == nil {
						account.state = newState
					}
				}
				account.nextAccessTime = time.Now().Add(2 * time.Second) // this time can be dynamicaly delayed
				acheap.Push(account)
			}
		}()
	}
	wg.Wait()
}
