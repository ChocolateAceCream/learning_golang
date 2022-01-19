package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("good")

	/*
	* Fprintln, Fprintf write content to an io.Write obj
	 */
	fmt.Fprintln(os.Stdout, "random context")

	file, err := os.OpenFile("1test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	time := time.Now()
	fmt.Fprintf(file, "info : %s\n", time)

	/*
	* Sprint, Sprintf write content to an string then return
	 */
	str := fmt.Sprintf("%2.2f", float64(1333/33))
	fmt.Println("str: ", str)
}
