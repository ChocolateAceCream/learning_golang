package Concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

//waitgroup is designed to sync multiple go routines together
var wg2 = sync.WaitGroup{}
var counter2 = 0

//  A read/write mutex allows all the readers to access the map at the same time, but a writer will lock out everyone else.
var m = sync.RWMutex{}

func rx() {
	//print out max available threads
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	//set available threads to 1
	runtime.GOMAXPROCS(1)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	//p.s. best practice is to set min threads to # of cores, and then tuning from there

	msg := "hello"
	wg2.Add(1)
	go func(msg string) {
		fmt.Println("msg from func2: ", msg)
		wg.Done()
	}(msg)
	wg2.Wait()
	msg = "bye"

	for i := 0; i < 10; i++ {
		wg2.Add(2)
		m.RLock()
		go sayNihao()
		m.Lock()
		go addOne()
	}
	wg2.Wait()
}

func sayNihao() {
	fmt.Println("from sayHello func: ", counter2)
	m.RUnlock()
	wg2.Done()
}

func addOne() {
	counter2++
	m.Unlock()
	wg.Done()
}
