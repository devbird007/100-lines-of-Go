package main

import (
	"fmt"
)

func foo(a int) {
	a = 10
}

func main() {
	var a int
	fmt.Printf("The value of a is %v\n", a)

	foo(a)
	fmt.Printf("The value of a is %v\n", a)
}
