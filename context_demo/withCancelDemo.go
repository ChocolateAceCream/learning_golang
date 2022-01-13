package main

import (
	"context"
	"fmt"
)

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return //end this goroutine, prevent memory leak

			// case ch <-n:
			// 	n++
			default:
				n++
				ch <- n
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println("count: ", n)
		if n > 5 {
			break
		}
	}
}
