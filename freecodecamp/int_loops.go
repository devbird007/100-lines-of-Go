package main

import (
	"fmt"
)

func main() {

	numbers := []int{1, 3, 5, 7, 9, 11}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
