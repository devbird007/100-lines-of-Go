package main

import (
	"fmt"
)

func main() {
	p := new(int)

	fmt.Printf("The data type of \"p\" is %T\n", p)
	fmt.Printf("Pointer p points to %v\n", *p)

	*p = 5
	fmt.Printf("Pointer p points to %v\n", *p)
}
