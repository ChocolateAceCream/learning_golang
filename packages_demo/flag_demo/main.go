/*
Usage:
1. go build
2. execute following command
./test -name=di -age=1 --gender male --flagname 33
*/
package main

import (
	"flag"
	"fmt"
)

//all of these three var are pointers
var cliName = flag.String("name", "nick", "Input your name")
var cliAge = flag.Int("age", 12, "Input your age")
var cliGender = flag.String("gender", "male", "Input your Gender")

// 定义一个值类型的命令行参数变量，在 Init() 函数中对其初始化
// 因此，命令行参数对应变量的定义和初始化是可以分开的
var cliFlag int

func Init() {
	flag.IntVar(&cliFlag, "flagname", 1234, "just for dmeo")
}

func main() {
	Init()

	//flag.Parse() need to be executed after all initialization of flags
	flag.Parse()

	// flag.Args() 函数返回没有被解析的命令行参数
	// func NArg() 函数返回没有被解析的命令行参数的个数
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i < flag.NArg(); i++ {
		fmt.Printf("args[%d]=%s\n", i, flag.Arg(i))
	}

	fmt.Println("name =", *cliName)
	fmt.Println("age =", *cliAge)
	fmt.Println("gender =", *cliGender)
	fmt.Println("flagname =", cliFlag)
}
