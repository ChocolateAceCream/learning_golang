package main

import "fmt"

func main() {
	if true {
		fmt.Println("TRUE")
	}

	if false {
		fmt.Println("false")
	} else if true {
		fmt.Println("true")
	}

	//common use of initializer in a if statement
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	if key, ok := m1["a"]; ok {
		fmt.Println("key dd exist, which is ", key)
	} else {
		fmt.Println("key dd does not exist")
	}
	//however, the variable key is only accessable within the if statement scope
	// fmt.Println(key) will return error.

	//comparison operators that works for numerical operators: >, <, !=, >=, <=
	//for string operators, use ==

	//logical operators: ||, &&, !
	fmt.Println("not operator !!true: ", !!true)
	fmt.Println("not operator !true: ", !true)
	bollean := returnTrue()
	fmt.Println("return value from bollean func: ", bollean)

	//short circuiting demo
	fmt.Println("----------short circuiting demo-------------")

	if 2 > 1 || returnTrue() {
		fmt.Println("or case short circuiting demo: ", bollean)
	}

	if 2 < 1 && returnTrue() {
		fmt.Println("and case short circuiting demo: ", bollean)
	}

	i := 2 > 1 && returnTrue()
	fmt.Println("i: ", i)

}

func returnTrue() bool {
	fmt.Println("return true function")
	return true
}
