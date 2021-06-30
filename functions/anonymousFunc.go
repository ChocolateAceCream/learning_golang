package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous function")
	}()

	for i := 0; i < 5; i++ {
		func() {
			//i is accessable insdie the anonymous func since it's defind out of the current scope.
			fmt.Println(i)
		}()

		//better approach is to pass in the i in anonymous func so it's scope safe when anonymous func is async
		func(i int) {
			fmt.Println(i)
		}(i)
	}
	//assign func() as a variable
	// alternative way of assign func() to variables:
	// var f func() = {}
	fmt.Println("--------assign func to variables------")
	f := func() {
		fmt.Println("Anonymous function")
	}
	f()

	//when using variables which contain a func(), make sure it's called after declaration
	var fn func(int, int) (int, error)
	fn = func(a, b int) (int, error) {
		sum := a + b
		return sum, nil
	}
	sum, err := fn(1, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sum)

}
