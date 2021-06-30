package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		subFunc("this is a msg", i)
	}
	paramTest("a", "b", "c", 1, 2)
	msg := "msg"
	ptr := "pointer"
	pointerTest(msg, &ptr)
	fmt.Printf("msg: %v; pointer: %v\n", msg, ptr)

	variadicParams("msg", 1, 2, 3, 4, 5)

	fmt.Println("----return function----")
	aa, bb := returnDemo()
	fmt.Println(aa, bb)

	fmt.Println("----function return pointer----")
	ptr2 := returnPtr()
	fmt.Println(*ptr2)

	fmt.Println("----function return pointer with syntax sugar----")
	ptr3 := syntaxSugar()
	fmt.Println(ptr3)

}

func subFunc(msg string, index int) {
	fmt.Printf("index: %v; value: %v\n", index, msg)
}

//for multiple params with same time, only need to separate them with ',' and append the type to the last param
func paramTest(a, b, c string, d, e int) {
	fmt.Println(a, b, c, d, e)
}

func pointerTest(msg string, ptr *string) {
	fmt.Println("---------------inside the pointer test function-------------------")
	msg = "new msg"
	*ptr = "new pointer"
	fmt.Println(*ptr, msg)
}

//variadic params can only be passed as the last params into the function
func variadicParams(msg string, values ...int) {
	fmt.Println("-------------variadicParams example, values: ", values)
	sum := 0
	for index, val := range values {
		fmt.Println(index)
		sum += val
	}
	fmt.Println("sum is: ", sum)
}

//function with multiple return values,
func returnDemo() (int, int) {
	return 123, 321
}

//function that return a pointer
// unlike other languages, which assign local stack of memories to execute the function (and GC  after execution),
// go will detect that you return a value generated on a local stack, it will use shared memory(heap) to store the variable
// so the value will not be affect by GC
func returnPtr() *int {
	a := 123
	return &a
}

// syntax sugar: assign a return variable in function declaration,
// so we don't need to declare return variable or initialize result
// don't use this when function is long, which will result in hard-to-read codes
// however, since return value is not declared explicitly,
// return a pointer will cause panic since GO didn't recognize the return value as pointer
// then no shared memory will be assgin to store the value
/*
	e.g. the following code will cause panic
	func syntaxSugar() (result *int) {
		// a := 123
		*result = 123
		return
	}
*/
func syntaxSugar() (result int) {
	// a := 123
	result = 123
	return
}
