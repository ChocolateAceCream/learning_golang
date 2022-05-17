package main

import "fmt"

func main() {
	greeting("a", "b", "c", "d")

	// be careful when passing slice to func with variadic params intake, because func greeting will not create new slice, but share the same reference of passed in slice. so if func greeting changed slice inside, the original slice will also be changed
	dd := []string{"b", "c", "d"}
	greeting("a", dd...)
	fmt.Println(dd)

}

func greeting(name1 string, who ...string) {
	if who == nil {
		fmt.Println("who is empty")
		return
	}
	who[1] = who[1] + "..."
	for _, name := range who {

		fmt.Println("Hi ", name)
	}

}
