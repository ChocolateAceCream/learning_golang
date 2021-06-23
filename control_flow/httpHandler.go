package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello go111!"))
	})
	err := http.ListenAndServe(":8080", nil)

	//use recover to capture error message and print to log
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	if err != nil {
		panic(err.Error())
	}

}
