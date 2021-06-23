package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			//in order not to let main func print out "end", we need to re-panic here
			// panic("repanic")
		}
	}()
	panic("something wrong")
	fmt.Println("done panicing") //won't be executed
}
