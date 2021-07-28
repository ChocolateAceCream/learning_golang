package main

import "fmt"

type StuffClient interface {
	DoStuff() error
}

func (s stuffClient) DoStuff() error {
	return nil
}

type stuffClient struct {
	timeout int
	retries int
}

type StuffClientOption func(*stuffClient)

func WithRetries(i int) StuffClientOption {
	//this anomynous function is of type StuffClientOption since it takes the same param *stuffClient as StuffClientOption
	return func(c *stuffClient) {
		c.retries = i
	}
}

func WithTimeout(i int) StuffClientOption {
	//this anomynous function is of type StuffClientOption since it takes the same param *stuffClient as StuffClientOption
	return func(c *stuffClient) {
		c.timeout = i
	}
}

//constructor
func NewStuffClient(opts ...StuffClientOption) StuffClient {
	client := stuffClient{
		retries: 3,
		timeout: 2,
	}
	for _, caller := range opts {
		caller(&client)
	}
	return client
}

func main() {
	x := NewStuffClient()
	fmt.Println(x) // prints &{{} 2 3}
	x = NewStuffClient(
		WithRetries(1),
	)
	fmt.Println(x) // prints &{{} 2 1}
	x = NewStuffClient(
		WithRetries(1),
		WithTimeout(1),
	)
	fmt.Println(x) // prints &{{} 1 1}
}
