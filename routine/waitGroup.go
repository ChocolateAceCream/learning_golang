package main

import (
	"fmt"
	"sync"
)

//waitgroup is designed to sync multiple go routines together
var wg = sync.WaitGroup{}
var counter = 0

func main() {
	msg := "hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println("msg from func2: ", msg)
		wg.Done()
	}(msg)
	wg.Wait()
	msg = "bye"

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go addOne()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Println("from sayHello func: ", counter)
	wg.Done()
}

func addOne() {
	counter++
	wg.Done()
}
