package main

import "fmt"

func main() {

	//-----------------------------array syntax-----------------------------
	grades := [3]float64{1, 2} //array can only store one type of data
	fmt.Printf("Grades: %v, %T, capacity: %v\n", grades, grades, cap(grades))

	age := [...]int{11, 22, 33} // ... means create an array so it can just hold all the variables in literal assignments
	fmt.Printf("age: %v\n", age)

	var students [3]string
	fmt.Printf("students: %v\n", students) //create an empty string array with length 3

	students[0] = "lisa"
	fmt.Printf("students: %v\n", students)
	fmt.Printf("students: %v\n", students[0])                // get the value of certain index in array
	fmt.Printf("students array length: %v\n", len(students)) // use build-in len() to retrive the length of array

	//------------------2D array--------------------
	fmt.Println("----------2d array-----------")

	var twoD [3][3]int
	twoD[0] = [3]int{1, 0, 0}
	twoD[1] = [3]int{0, 1, 0}
	twoD[2] = [3]int{0, 0, 1}
	fmt.Printf("twoD: %v, %T\n", twoD, twoD)

	//copy array will not pass reference, go pass the data by default
	twoD2 := twoD
	twoD2[0][0] = 3
	fmt.Printf("twoD2: %v\n", twoD2)

	//in order to pass by reference, we use pointer
	twoD3 := &twoD
	twoD3[1] = [3]int{4, 4, 4}
	fmt.Printf("twoD: %v\n", twoD2)
	fmt.Printf("twoD3: %v\n", twoD2)

	//-----------------------------slice syntax-----------------------------
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("slice: %v, length: %v, capacity: %v\n", a, len(a), cap(a))
	b := a[:]   //[0 1 2 3 4 5 6 7 8 9 10]
	c := a[:4]  //[0 1 2 3]
	d := a[5:]  //[5 6 7 8 9 10]
	e := a[3:8] //[3 4 5 6 7]
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %v\n", e)

	fmt.Println("--------change b[1] to 100--------------")
	//since slice pass by reference, we change on one location, all get effected
	b[1] = 100
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %v\n", e)

	//slice operations also works on arrays, but it will affect original array because new slice refer to the address of original array
	fmt.Println("--------change b[2] to 100--------------")
	b[2] = 100
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %v\n", e)

	//use make() to create slice
	make := make([]int, 3)
	fmt.Println(make)                       //[0,0,0]
	fmt.Printf("length: %v\n", len(make))   //by default, slice was filled with 0, so length is 0
	fmt.Printf("capacity: %v\n", cap(make)) // capacity is also 0

	/*
		append function (used to push element to slice): always use its return value to assign back to original variables.
		Reason: when bottom array of slice object is full, append() will return a new slice, so always remember to use the return value to re-assign the original slice object
		p.s. append nil to slice with append() will not cause an error, so be careful when append nil to slice
	*/
	make = append(make, 12, 3, 4, 5) //append() can take multiple arguments, first is the slice it append to, the rest arguments are the value to be append
	fmt.Println(make)
	fmt.Printf("length: %v\n", len(make))
	fmt.Printf("capacity: %v\n", cap(make))
	//append() can only take arguments value with same type of existing elements.
	//e.g. 2D slice can only append 1D slice to it

	t := [][]int{{1, 2}, {3, 4, 5}}           //create 2d slice
	t = append(t, [][]int{{4, 4}, {5, 5}}...) // use spread operator
	fmt.Printf("%v, %T\n", t, t)

	//shift operation on slice
	leftShift := t[1:] //remove fist element of slice
	fmt.Printf("t: %v, %T\n", t, t)
	fmt.Printf("left shift t: %v, %T\n", leftShift, leftShift)
	rightShift := t[:len(t)-1] //remove last element of slice
	fmt.Printf("right shift t: %v, %T\n", rightShift, rightShift)
	middle := append(t[:2], t[3:]...) // be careful when you do this operation, since it will also change the original array.
	fmt.Printf("middle: %v, %T\n", middle, middle)
	fmt.Printf("t: %v, %T\n", t, t)

}
