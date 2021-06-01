package main

import (
	"fmt"
)

const (
	isGuest = 1 << iota
	isAdmin
	isDeveloper
	isBoss
)

func main() {
	var role byte = isAdmin | isBoss
	fmt.Printf("%b\n", role)                              // print out byte value of role
	fmt.Printf("is admin? %v\n", isAdmin&role == isAdmin) // check if it's admin
	fmt.Printf("is guest? %v\n", isGuest&role == isGuest)
	role = role | isGuest //assign guest to role
	fmt.Printf("is guest? %v\n", isGuest&role == isGuest)
	role &^= isGuest // use clear bit to unset role guest
	fmt.Printf("is guest? %v\n", isGuest&role == isGuest)
}
