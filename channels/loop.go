package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	ch := make(chan int, 50)
	go func(ch chan<- int) {
		ch <- 21
		ch <- 31

		/*
			- close channel so range function knows new more msgs are coming
			- any funcs that have access to the channel can trigger the close func
			- once channel closed, you cannot pass msg to that channel
			- you cannot reopen a closed channel
			- you cannot detect if channel is closed, except for looking for panic
		*/
		close(ch)
		wg.Done()
	}(ch)
	go func(ch <-chan int) {
		//slice over channel is different, (unlike arr etc, first e is index and second e is value)
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	//alternative way: use for loop
	ch2 := make(chan int, 50)
	wg.Add(2)
	go func(ch2 <-chan int) {
		for {
			if i, ok := <-ch2; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch2)
	go func(ch2 chan<- int) {
		ch2 <- 4
		ch2 <- 7
		close(ch2)
		wg.Done()
	}(ch2)
	wg.Wait()
}
