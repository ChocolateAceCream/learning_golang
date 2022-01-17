package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Key string

var wg = sync.WaitGroup{}

func worker(ctx context.Context) {
	key := Key("keyName")
	value := ctx.Value(key).(string)
	fmt.Println("value: ", value)
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker done")
			break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	key := Key("keyName")
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	ctx = context.WithValue(ctx, key, "actual value")
	wg.Add(1)
	go worker(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
}
