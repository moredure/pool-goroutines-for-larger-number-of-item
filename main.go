package main

import (
	"log"
	"time"
)

var acheap = &SyncMutexItemsHeap{}

var accounts Accounts

var N = 10

func main() {
	// listen for nats on new accounts
	// listen for nats on remove account
	for i := 0; i < N; i += 1 { // N is number of workers
		// can be dynamicaly resized if size of acheap grows?
		go func() {
			for {
				account := acheap.Pop()
				if account.nextAccessTime.After(time.Now()) {
					acheap.Push(account)
					continue
				}
				select {
				case <-account.stopped:
					continue
				default:
					var err error
					var newState int
					switch account.state {
					case 0:
						// in case of error retry
						// in case of result return next state
						newState, err = account.SomeAction()
					case 1:
						newState, err = account.SomeAction2()
					default:
						log.Fatal("undefined state")
					}
					if err == nil {
						account.state = newState
					}
				}
				account.nextAccessTime = time.Now().Add(5 * time.Second) // this time can be dynamicaly delayed
				acheap.Push(account)
			}
		}()
	}
}
