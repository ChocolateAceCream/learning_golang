package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	TickerLaunch()
}

func TickerLaunch() {
	ticker := time.NewTicker(1 * time.Second)
	//always remember to defer ticker stop in order to prevent resource leaking
	defer ticker.Stop()
	max := 10
	queue := []int{}

	for {
		randIntn := rand.Intn(20)
		select {
		case <-ticker.C:
			queue = append(queue, randIntn)
			fmt.Println(queue)
		default:
			if len(queue) > max {
				fmt.Println("done")
				return
			}
		}
	}
}
