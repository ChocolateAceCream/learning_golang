package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	f, err := os.ReadFile("sample.csv")
	if err != nil {
		log.Fatal(err)
	}
	println(string(f))
	fileOpen()
}

func fileOpen() {
	file, err := os.Open("sample.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	count := 0
	header, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("header: ", header)
	arr := [][]string{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(count, record)
		arr = append(arr, record)
		count++
	}
	fmt.Println("initial: ", arr)
	fmt.Println("				: ", arr[0])

	options := []int{}

	for i := 1; i < len(arr); i++ {
		options = append(options, i)
	}
	fmt.Println("options: ", options)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	next := 0
	lastElement := arr[0]

	for len(options) >= 1 {
		randIdx := rnd.Intn(len(options))
		for options[randIdx] == next {
			randIdx = rnd.Intn(len(options))
		}
		fmt.Println("rand idx:", options[randIdx])

		arr[next] = arr[options[randIdx]]
		next = options[randIdx]
		options = append(options[:randIdx], options[randIdx+1:]...)
		fmt.Println(arr)
		fmt.Println("options: ", options)
	}
	fmt.Println("next: ", next)
	arr[next] = lastElement
	fmt.Println("result:", arr)
}
