package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// var repeat = flag.Int("r", 10000, "Input repeat times")
var count = 0
var repeat = 1000000
var wg = sync.WaitGroup{}

func randGenerator() {
	rand.Seed(time.Now().UnixNano())
	rand1 := rand.Float64() / 2

	rand2 := rand.Float64()

	// fmt.Printf("rand1: %v; rand2: %v\n", rand1, rand2)
	n1 := rand1
	n2 := (1 - rand1) * rand2
	n3 := (1 - rand1) * (1 - rand2)
	// fmt.Println("sum: ", n1+n2+n3)
	arr := []float64{n1, n2, n3}
	sort.Float64s(arr)
	// fmt.Printf("arr: %v\n", arr)
	result := arr[0]+arr[1] > arr[2]
	// fmt.Println("---result---", result)
	if result {
		count++
	}
}

func SingleThread() {
	count = 0
	for i := 0; i < repeat; i++ {
		randGenerator()
	}
	fmt.Printf("singleThread count: %v\n", count)
}

func MultiThread() {
	count = 0
	wg.Add(4)
	for j := 0; j < 4; j++ {
		go func() {
			for i := 0; i < repeat/4; i++ {
				randGenerator()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("multiThread count: %v\n", count)
}
func main() {
	fmt.Printf("repeat %v times\n", repeat)
}
