/*
	simulate db connection, set default timeout to 10s, then create a ctx with timeout 2s (simulate db connetion time cost)
	so program ends when 5s reached
*/
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func worker(ctx context.Context) {
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("connection done")
			break LOOP
		default:
			fmt.Println("connecting...")
			time.Sleep(time.Second)
		}
	}
	wg.Done()
}

func main() {
	// if set timeout to 15 sec, then cancel() from main() will execute first
	// ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
}
