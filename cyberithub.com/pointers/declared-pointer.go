package main

import (
	"fmt"
)

func main() {
	var p *int

	fmt.Printf("The value stored in p is %v\n", p)
	fmt.Printf("The pointer p points to %v\n", *p)
}

// This file returns an error btw
