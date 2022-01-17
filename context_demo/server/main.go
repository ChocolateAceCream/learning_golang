package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

//simulate slow response

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// rand.Intn(2) return a positive int with range 0 ~ 2
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(10 * time.Second)
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprintf(w, "quick response")
}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
