/*
* @fileName factorial.go
* @author Di Sheng
* @date 2022/07/01 20:57:28
* @description: function to calculate factorial
 */

package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(factorial(4))
}

func factorial(n int64) *big.Int {
	fac := new(big.Int)
	fac.MulRange(1, n)
	return fac
}
