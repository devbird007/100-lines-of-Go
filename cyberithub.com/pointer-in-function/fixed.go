package main

import (
	"fmt"
)

func foo(a *int) {
	*a = 10
}

func main() {
	var a int
	fmt.Printf("The value of a is %v\n", a)

	p := &a
	foo(p)
	fmt.Printf("The value of a is %v\n", a)
}
