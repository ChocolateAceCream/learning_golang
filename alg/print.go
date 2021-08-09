package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	letter, number := make(chan bool), make(chan bool)
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
				// default:
				// 	break
			}
		}
	}()
	wg.Add(1)
	number <- true
	go func(wg *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVW"
		i := 0
		for {
			select {
			case <-letter:
				if i >= strings.Count(str, "")-1 {
					close(number)
					close(letter)
					wg.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				if i >= strings.Count(str, "")-1 {
					fmt.Println("")
					close(number)
					close(letter)
					wg.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				number <- true
				break
			default:
				break
			}
		}
	}(&wg)
	wg.Wait()
}
