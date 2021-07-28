//since function is a use-defined type, you can use functions as a field in a struct

package main

import "fmt"

type HelloFunc func()

/*
A function type denotes the set of all functions with the same parameter and result types.
since english() and chinese() bot have same params and return type with BycFunc, they can be convert to ByeFunc type
.e.g
	c := ByeFunc(english)
*/
type ByeFunc func(name string) string

func english(name string) string {
	return "bye, " + name
}

func chinese(name string) string {
	return "88, " + name
}

//append say() method to ByeFunc
func (g ByeFunc) say(name string) string {
	return g(name)
}

type Demo struct {
	bye   ByeFunc
	hello HelloFunc
}

func main() {
	a := Demo{
		hello: func() {
			fmt.Println("hello")
		},
	}
	a.hello()

	b := Demo{
		bye: func(name string) string {
			return "bye, " + name
		},
	}
	fmt.Println(b.bye("b"))

	//convert english to ByeFunc type, so we can call the say() method of ByeFunc type
	c := ByeFunc(english)
	fmt.Println(c.say("english"))
	c = ByeFunc(chinese)
	fmt.Println(c.say("chinese"))
}
