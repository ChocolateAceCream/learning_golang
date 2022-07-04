/*
* @fileName regexDemo.go
* @author Di Sheng
* @date 2022/07/04 10:46:35
* @description
 */

package main

import (
	"fmt"
	"regexp"
)
((e|E)[+-]?[0-9]+)?
func main() {
	r, _ := regexp.Compile(`^[+-]?([0-9]+\.?[0-9]*((e|E)[+-]?[0-9]+)?|\.[0-9]+((e|E)[+-]?[0-9]+)?|)$`)
	fmt.Println(r.MatchString("e"))

	r2, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r2.MatchString("peach"))
}
