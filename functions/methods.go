package main

import "fmt"

func main() {
	g := greeter{
		name:     "John",
		greeting: "hello",
	}
	g.method()
	fmt.Println("new name is : ", g.name)
	g.change()
	fmt.Println("after change, name is : ", g.name)

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

// greeter is value receiver, so method will be append to a copy of greeter struct
//everytime we invoke this method, we will create a copy of that greeter struc (which is bad when greeter is a large struct)
//but the benifits are that the parent struct won't be manipulated inside the method
func (g greeter) method() {
	fmt.Println(g.greeting, g.name)
	g.name = "new name"
}

func (i plate) count(num plate) plate {
	return i + num
}

//alternatively, we can pass the pointer as a value receiver, in which case we will be able to manipulated
// the parent struct inside the method
func (i *greeter) change() {
	i.name = "hey hey hey"
}
