package main

import (
	"sync"
	"fmt"
)

var ledger int = 0
// var m sync.Mutex
var ch chan int
var wg sync.WaitGroup

func worker1(){
 	// m.Lock()
	ch<-1
 	// ledger = 1
 	// m.Unlock()
	// <-ch
	wg.Done()
}

func worker2(){
 	// m.Lock()
 	ch<-2
	// ledger = 2
 	// m.Unlock()
	// <-ch
	wg.Done()
}

func worker3(){
	// m.Lock()
	ch <- 3
	// ledger = 3
	// <- ch
	// m.Unlock()
	wg.Done()
}

func main(){
	ch = make(chan int)
	go func() {
		for {
			select {
			case request := <-ch:
				ledger = request
			}
		}
	}()
	wg.Add(3)
	go worker1()
   	go worker2()
   	go worker3()
	wg.Wait()
	fmt.Printf("%s\n", "completed")
}