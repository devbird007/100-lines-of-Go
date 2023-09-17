package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	var a = 10
	var b = 15

	// Function call where we pass in the two variables a and b
	// and store the result in c
	var c = add(a, b)
	fmt.Println("The sum of a and b is ", c)
}
