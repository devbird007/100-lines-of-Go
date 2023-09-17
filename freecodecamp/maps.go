package main

import (
	"fmt"
)

func main() {
	books := map[string]int{
		"maths":     5,
		"biology":   9,
		"chemistry": 6,
		"physics":   3,
	}
	for _, val := range books {
		fmt.Println(val)
	}
}
