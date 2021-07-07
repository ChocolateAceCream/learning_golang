package main

import "fmt"

//定义接口R、W、RW
type R interface {
	Read()
}
type W interface {
	Write()
}
type RW interface {
	R
	W
}

//定义类型Vertex，实现R接口
type Vertex struct {
	x int
}

func (v Vertex) Read() {
	fmt.Println("fun read")
}

func (v Vertex) Write() {
	fmt.Println("fun write")
}

//func (v Vertex) Write(){
//	fmt.Println("fun")
//}
func main() {
	var a = Vertex{x: 12345}
	var i RW = &a //不报错，&a实现了R接口

	r, rok := i.(R)            //check if i implement R interface
	fmt.Println(r, rok)        //-->&{} true
	fmt.Println(r.(*Vertex).x) //-->&{} true
	w, wok := i.(W)
	fmt.Println(w, wok) //--><nil> false
	rw, rwok := i.(RW)
	fmt.Println(rw, rwok) //--><nil> false

	do(21)
	do("hello")
	do(true)
	do(Vertex{})
}

func do(i interface{}) { //形参是一个空接口
	switch v := i.(type) { //注意type关键字
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
