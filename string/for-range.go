package main

import "fmt"

func main() {

	// string,
	s := "hello你好"

	// range
	for index, runeValue := range s {
		fmt.Printf("%q starts at byte position %v\n", runeValue, index)

	}
}
