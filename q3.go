package main

// import "sync"

var ledger int = 0
// var m sync.Mutex
var ch chan int

func worker1(){
 	// m.Lock()
	ch<-1
 	ledger = 1
 	// m.Unlock()
	<-ch
}

func worker2(){
 	// m.Lock()
 	ch<-1
	ledger = 2
 	// m.Unlock()
	<-ch
}

func worker3(){
	// m.Lock()
	ch <- 1
	ledger = 3
	<- ch
	// m.Unlock()
}

func main(){
	ch = make(chan int)
	go worker1()
   	go worker2()
   	go worker3()
}