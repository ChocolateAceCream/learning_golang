package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	ch := make(chan int)

	//receive-only channel: can only write to the channel, read from channel will cause error
	go func(ch chan<- int) {
		ch <- 4
		wg.Done()
	}(ch)

	//send-only channel: can only read from channel, write to the channel will cause error
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	wg.Wait()
}
