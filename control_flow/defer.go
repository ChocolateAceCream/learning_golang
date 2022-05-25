package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	actual flow of main() function is that
	fmt.Println("first")
	then it sees defer, so jump to
	fmt.Println("third")
	after execution of all its statement done, right before function return(exit)
	go will check for any defer statements and call it in Last in first out order .
*/
func main() {
	fmt.Println("first")
	defer fmt.Println("defer first")
	fmt.Println("second")
	defer fmt.Println("defer second")
	fmt.Println("third")
	defer fmt.Println("defer third")

	fmt.Println("----variable changed example---------")
	a := 10
	//will print a= 10 because the when you defer a function, the args of function are taken at the time defer is called,
	// not the time defer is executed.
	defer fmt.Println("a = ? ", a)
	a = 20

	res, err := http.Get("https://gitlab.ctbiyi.com/smart_manage/smart_important_window_mobile/commits/dev")
	if err != nil {
		log.Fatal(err)
	}
	//this most common useage of defer: closing http request is preferred so no need to worry about other logic in between
	//open and close request, since this defer statement will be executed last
	//p.s. when open resource in loop, better avoid using this pattern
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	fmt.Println("------DeferFunc()---start----")
	result := DeferFunc()
	fmt.Println(result)
	fmt.Println("------DeferFunc()-end------")
}

/*
actual execution order is that:
1. result = i
2. result ++
3. return
*/
func DeferFunc() (result int) {
	i := 1
	defer func() {
		result++
	}()
	return i
}
