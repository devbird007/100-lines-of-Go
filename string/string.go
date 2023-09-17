package main

import "fmt"

func main() {
	// string is a slice of bytes.
	s := "hello你好"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x\n", s[i])
	}
}
