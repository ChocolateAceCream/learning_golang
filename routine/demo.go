/*
# Best practices

- Don't create go routine in lib, let consumer control concurrency
- when create go routine, remember to end it (to avoid memory leaks)
- check for race conditions at compile time by using go run -race
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

//waitgroup is designed to sync multiple go routines together
var wg = sync.WaitGroup{}

func main() {
	//let go create a green thread and run sayHello() func in that thread
	msg := "hello"

	//anonymous func can access to variables from outer scope, even though go routine is running on seperated execution stack
	//however, this will print because go schedular will first execute the main func() before any go routines
	//P.S, that's why we append time.Sleep() func call at the end of main func() in order to wait for go routines to be executed
	//so this implementation will cause a race condition which we should avoid
	go func() {
		fmt.Println("msg from func1: ", msg)
	}()

	//best practice is to decouple the go routine and main func() by passing msg params
	go func(msg string) {
		fmt.Println("msg from func2: ", msg)
	}(msg)
	go sayHello(msg)
	msg = "bye"
	time.Sleep(5000 * time.Millisecond)
}

func sayHello(msg string) {
	fmt.Printf("from sayHello: %v\n", msg)

}
