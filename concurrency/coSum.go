package Concurrency

import (
	"math/rand"
	"sync"
	"time"
)

var wg5 sync.WaitGroup

func CoSum(total int, threads int) int {
	cbChannel := make(chan int, threads)
	incomeChan := make(chan int, threads)
	for i := 0; i < threads; i++ {
		wg5.Add(1)
		go sumCalculator(cbChannel, incomeChan)
	}
	for i := 0; i < total; i++ {
		j := i
		incomeChan <- j
	}
	close(incomeChan)
	sum := 0
	r := 0
	wg5.Wait()
	for {
		if partialSum, ok := <-cbChannel; ok {
			r++
			sum += partialSum
		}
		if r == threads {
			close(cbChannel)
			break
		}
	}
	return sum
}

func sumCalculator(cbChannel chan int, incomeChan chan int) {
	rand.Seed(time.Now().UnixNano())
	sum := 0
	for i := range incomeChan {
		time.Sleep(1000)
		sum += i
	}
	cbChannel <- sum
	wg5.Done()
}

func singleTreadSum(total int) int {
	sum := 0
	for i := 0; i < total; i++ {
		time.Sleep(1000)
		sum += i
	}
	return sum
}
