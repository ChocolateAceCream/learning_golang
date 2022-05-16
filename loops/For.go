package main

import (
	"fmt"
)

func main() {
	// initialize counter this way, then counter is only accessable inside the for loop scope.
	for i := 0; i < 5; i = i + 2 {
		fmt.Println(i)
	}
	// counter could also be initialized else where, then it will be accessible outside the for loop
	//the incrementor can only take statement, and it's not required(you can increment counter insdie loop explicitly)
	//then the for loop will turn to a while loop
	i := 0
	fmt.Println("----i----")
	for {
		i++
		if i == 3 {
			fmt.Println("----return---")
			return
		}
	}
	fmt.Println("-------i:", i)

	//break is used to jump out of the nearest loop(only that loop), in this case, the infinite loop
	//if that loop is wrapped by other loops, then other loops will continue their iteration
	for {
		i = i + 4
		if i > 100 {
			break //jump out of the loop
		}
	}

	//optional, we can use label Loop:  to tell go where we want to break the loop to
	fmt.Println("---Loop break example----")
Loop:
	for i := 0; i < 5; i++ {
		// Loop: also can be placed here
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				fmt.Println("---i, j,k----", i, j, k)
				if k > 3 {
					break Loop
				}
			}
		}
	}

	//use continue to skip current iteration
	fmt.Println("-----continue example---")
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("infinite loop", i)

	//looping two variables
	fmt.Println("looping two variables")
	for i, j := 1, 2; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	//use for loop to iterate a collection, such as slice, map, string etc..
	fmt.Println("------collection iteration with for loop------------")
	slice := []int{111, 222, 333, 444, 555}
	for i, v := range slice {
		fmt.Printf("value: %v, index: %v \n", v, i)
	}

	fmt.Println("------loop through string------------")
	str := "asdfasdf"
	for i, v := range str {
		// value will be returned in form of unicode by default, so we need to cast it to string
		fmt.Printf("value: %v, index: %v \n", string(v), i)
	}
	fmt.Println("------loop through map------------")
	m := map[string]int{
		"h": 1,
		"e": 2,
	}

	//if only value is required : for _, v := range m
	//if only key is required : for k := range m
	for k, v := range m {
		fmt.Printf("value: %v, key: %v \n", v, k)
	}

}

/*
	P.S: the fact is: in each iteration of
		for index, value := range xxx
	index and value are re-assigned to new value (not re-created ). so take extra care when including index and value in  each iterations (especially in goroutine)
*/
// wrong example : the taks printed in goroutine is randomly, since task is bonded by its reference not its value
func Process1(tasks []string) {
	for _, task := range tasks {
		// 启动协程并发处理任务
		go func() {
			fmt.Printf("Worker start process task: %s\n", task)
		}()
	}
}

// good example: bonding value, not reference
func Process2(tasks []string) {
	for _, task := range tasks {
		go func(t string) {
			fmt.Printf("Worker start process task: %s\n", t)
		}(task)
	}
}

type Tests struct {
	name         string
	input        int
	expectOutput int
}

var tests = []struct {
	name         string
	input        int
	expectOutput int
}{}

func Proces3() {
	t := []Tests{
		{
			name:         "double 1 should got 2",
			input:        1,
			expectOutput: 2,
		},
		{
			name:         "double 2 should got 4",
			input:        2,
			expectOutput: 4,
		},
	}

	for _, test := range t {
		//fix: declare another variable to replace loop variable explicitly
		ts := test

		go func(name string) {
			fmt.Printf("Worker start process task: %s\n", name)

			//name is bonded with test.value, however, input is not, so this may cause a problem
			// fmt.Printf("Worker start process task: %c\n", test.input)
			fmt.Printf("Worker start process task: %c\n", ts.input)

		}(test.name)
	}
}
