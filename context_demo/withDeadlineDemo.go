package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(5*time.Second))

	//a good practice is defer cancel() so no matter what ctx will be cancelled eventually
	defer cancel()
	count := 0

	for {
		select {
		// time.After(d duration): first elapse d duration then send the current time on the return channel
		case <-time.After(1 * time.Second):
			fmt.Println("count: ", time.Now())
			count++
		case <-ctx.Done():
			fmt.Println("done: ", count)
			return
		}
	}
}
