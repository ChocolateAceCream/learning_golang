package main

import "fmt"

func main() {
	var pw PlateWritter = Car{}
	pw.WritePlate([]byte("mazda"))

	var emptyInterface interface{} = pw
	emptyInterface.(PlateWritter).WritePlate([]byte("emptyInterface"))

}

//interface describe behaviors, store method definitions
type PlateWritter interface {
	WritePlate([]byte) (int, error) //a method take []bytes as input and return int and error
}

type Car struct{}

func (car Car) WritePlate(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
