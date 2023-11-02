package Concurrency

import (
	"fmt"
	"sync"
)

//waitgroup is designed to sync multiple go routines together
var wg1 = sync.WaitGroup{}
var counter1 = 0

func wgFunc() {
	msg := "hello"
	wg1.Add(1)
	go func(msg string) {
		fmt.Println("msg from func2: ", msg)
		wg1.Done()
	}(msg)
	wg1.Wait()
	msg = "bye"

	for i := 0; i < 10; i++ {
		wg1.Add(2)
		go sayHi()
		go add1()
	}
	wg1.Wait()
}

func sayHi() {
	fmt.Println("from sayHello func: ", counter1)
	wg.Done()
}

func add1() {
	counter1++
	wg1.Done()
}
