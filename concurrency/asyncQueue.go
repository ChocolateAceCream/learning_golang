package Concurrency

import (
	"fmt"
	"sync"
	"time"
)

const MAXOUTSTANDING = 5

func handle(queue chan int, wg4 *sync.WaitGroup, j int) {
	defer wg4.Done()
	for r := range queue {
		time.Sleep(2000)
		fmt.Printf("process #%v: %v\n", j, r)
	}
}

func Serve(clientRequests chan int, wg4 *sync.WaitGroup) {
	// Start handlers
	for i := 0; i < 5; i++ {
		wg4.Add(1)
		j := i
		go handle(clientRequests, wg4, j)
	}
	fmt.Println("before quit-----")
}

func MyFunction() {
	var wg4 sync.WaitGroup

	clientRequests := make(chan int, 100)
	Serve(clientRequests, &wg4)
	for i := 0; i < 20; i++ {
		clientRequests <- i
	}
	close(clientRequests)
	fmt.Println("end")
	wg4.Wait()
}
