package main

import "fmt"

func main() {
	g := greeter{
		name:     "John",
		greeting: "hello",
	}
	g.method()
	fmt.Println("new name is : ", g.name)

	var a plate = 1
	b := 2
	c := a.count(plate(b))
	fmt.Println(c)
}

type greeter struct {
	greeting string
	name     string
}

type plate int

func (g greeter) method() {
	fmt.Println(g.greeting, g.name)
	g.name = "new name"
}

func (i plate) count(num plate) plate {
	return i + num
}
