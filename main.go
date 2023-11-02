package main

import (
	"fmt"
	"sync"
	"time"
)

const MAXOUTSTANDING = 5

func main() {
	var wg sync.WaitGroup

	clientRequests := make(chan int, 100)
	Serve(clientRequests, &wg)
	for i := 0; i < 20; i++ {
		clientRequests <- i
	}
	close(clientRequests)
	fmt.Println("end")
	wg.Wait()
}

func handle(queue chan int, wg *sync.WaitGroup, j int) {
	defer wg.Done()
	for r := range queue {
		time.Sleep(2000)
		fmt.Printf("process #%v: %v\n", j, r)
	}
}

func Serve(clientRequests chan int, wg *sync.WaitGroup) {
	// Start handlers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		j := i
		go handle(clientRequests, wg, j)
	}
	fmt.Println("before quit-----")
}
