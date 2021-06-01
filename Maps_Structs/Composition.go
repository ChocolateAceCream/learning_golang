package main

import "fmt"

func main() {
	type P struct {
		a int
		b int
	}
	type F struct {
		P
		c int
		d int
	}

	p := P{
		a: 1,
		b: 2,
	}

	//structural assignment
	f := F{}
	f.a = 1
	f.b = 2
	f.c = 3
	f.d = 4
	fmt.Println("p", p)
	fmt.Println("structural assignment: ", f)

	//literal assignment
	g := F{
		P: P{a: 1, b: 2},
		c: 3,
		d: 4,
	}
	fmt.Println("literal assignment: ", g)

}
