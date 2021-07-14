package main

import (
	"fmt"
	"runtime"
	"sync"
)

//waitgroup is designed to sync multiple go routines together
var wg = sync.WaitGroup{}
var counter = 0

//  A read/write mutex allows all the readers to access the map at the same time, but a writer will lock out everyone else.
var m = sync.RWMutex{}

func main() {
	//print out max available threads
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	//set available threads to 1
	runtime.GOMAXPROCS(1)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	//p.s. best practice is to set min threads to # of cores, and then tuning from there

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
		m.RLock()
		go sayHello()
		m.Lock()
		go addOne()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Println("from sayHello func: ", counter)
	m.RUnlock()
	wg.Done()
}

func addOne() {
	counter++
	m.Unlock()
	wg.Done()
}
