//since function is a use-defined type, you can use functions as a field in a struct

package main

import "fmt"

type HelloFunc func()
type ByeFunc func() string

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
		bye: func() string {
			return "bye"
		},
	}
	fmt.Println(b.bye())
}
