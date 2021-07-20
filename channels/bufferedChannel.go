package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)

	//second param refers to buffered channel
	//50 means channel will have an internal store of size 50(not including the default, so channel can hold 51 maximun)
	ch := make(chan int, 50)

	go func(ch chan<- int) {
		ch <- 4
		ch <- 7
		wg.Done()
	}(ch)

	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	wg.Wait()
}
