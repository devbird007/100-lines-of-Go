package main

import "fmt"

func main() {
	var p *int
	var a int = 8

	// Assign the address of "a" to the pointer "p"
	p = &a

	fmt.Printf("The address of a is: %v\n", &a)
	fmt.Printf("The value stored in p is %v\n", p)
}
