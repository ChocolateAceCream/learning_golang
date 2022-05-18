package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	r := true
	var result *bool = &r
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	go WaitChannel(ch, result)
	for {
		select {
		case <-ticker.C:
			println("tick")
			ch <- time.Now().String()
		default:
			if *result == false {
				println("stop")
				return
			}
		}
	}
}

func WaitChannel(conn <-chan string, result *bool) {
	timer := time.NewTimer(5 * time.Second)
	for {
		select {
		case t := <-conn:
			// timer.Stop()
			fmt.Println(t)
			*result = true
		case <-timer.C:
			println("timeout")
			*result = false
		}
	}

}
