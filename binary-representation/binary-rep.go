package main

import "fmt"

func main() {
	for x := 0; x < 16; x++ {
		fmt.Printf("%b\n", x)
	}

	fmt.Println("Fixed length of 8 digits:")
	for x := 0; x < 16; x++ {
		fmt.Printf("\n %08b %#x %04d %v \n", x, x, x, x)
	}
}
