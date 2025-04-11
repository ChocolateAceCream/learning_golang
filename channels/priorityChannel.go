package main

import (
	"fmt"
	"time"
)

func main() {
	highPriority := make(chan int, 12)
	lowPriority := make(chan int, 12)
	highPriority <- 1
	lowPriority <- 2
	highPriority <- 3
	lowPriority <- 4
	for {
		select {
		case hp := <-highPriority:
			time.Sleep(time.Duration(hp) * time.Second)
			fmt.Println("High Priority Task Completed")
		default:
			select {
			case lp := <-lowPriority:
				time.Sleep(time.Duration(lp) * time.Second)
				fmt.Println("Low Priority Task Completed")
			}
		}
	}

}
