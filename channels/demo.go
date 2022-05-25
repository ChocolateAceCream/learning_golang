package main

import (
	"fmt"
	"sync"
)

//waitgroup is designed to sync multiple go routines together
var wg = sync.WaitGroup{}

func main() {
	ChannelIterator()

	//channel can only be created with make(), and can only send one type of data to the channel
	//unbuffered channel can only hold one data(in this case, one int variable)
	//push to a occupied channel will result in deadlock
	ch := make(chan int)
	wg.Add(5)
	go func() {
		i := <-ch
		fmt.Println("func 1.1: ", i)
		i = <-ch
		fmt.Println("func 1.2: ", i)
		wg.Done()
	}()
	go func() {
		i := <-ch
		fmt.Println("func 2.1: ", i)
		i = <-ch
		fmt.Println("func 2.2: ", i)
		wg.Done()
	}()
	go func() {
		i := <-ch
		fmt.Println("func 3: ", i)
		wg.Done()
	}()
	go func() {
		ch <- 1 //passing data into channel(pass the data copy)
		ch <- 2
		ch <- 3
		wg.Done()
	}()
	go func() {
		ch <- 4
		ch <- 5
		wg.Done()
	}()
	wg.Wait()

	//part two
	ch2 := make(chan int)
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go func() {
			ch2 <- i
			wg.Done()
		}()
		go func() {
			j := <-ch2
			fmt.Printf("index: %v, value: %v\n", i, j)
			wg.Done()
		}()
	}
	wg.Wait()

}

func ChannelIterator() {
	ch := make(chan int, 10)
	ch <- 4
	ch <- 1
	ch <- 2
	/*
		- close channel so range function knows new more msgs are coming
		- any funcs that have access to the channel can trigger the close func
		- once channel closed, you cannot pass msg to that channel
		- you cannot reopen a closed channel
		- you cannot detect if channel is closed, except for looking for panic
	*/
	// close(ch)

	// better separate ch write and read operations in different functions

	go func() {
		for val := range ch {
			fmt.Println(val)
		}
	}()

	fmt.Println("-----ChannelIterator done-----")
}
