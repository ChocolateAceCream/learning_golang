package main

import (
	"connectionpool/pool"
	"fmt"
	"time"
)

var idCounter int

func ConnectionFactory() (*pool.Connection, error) {
	idCounter++
	fmt.Println("idCounter: ", idCounter)
	return &pool.Connection{ID: idCounter}, nil
}

func main() {
	fmt.Println("----main-----------")
	pool, err := pool.NewPool(2, ConnectionFactory)
	if err != nil {
		fmt.Println("Error creating pool:", err)
		return
	}

	// Simulate getting and releasing connections.
	for i := 0; i < 5; i++ {
		go func(i int) {
			conn, err := pool.Get()
			if err != nil {
				fmt.Println("Error getting connection:", err)
				return
			}
			fmt.Printf("Goroutine %d: Using connection %d\n", i, conn.ID)
			time.Sleep(1 * time.Second) // Simulate work
			pool.Release(conn)
		}(i)
	}

	// Wait for a moment to allow goroutines to finish
	time.Sleep(13 * time.Second)

	// Close the pool
	pool.Close()
}
