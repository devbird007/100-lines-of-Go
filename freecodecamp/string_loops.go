package main

import (
	"fmt"
)

func main() {
	word := "Ab$du"

	for index, a := range word {
		fmt.Println(index, string(a))
	}

	table := "Viseri0N#"

	for i := 0; i < len(table); i++ {
		fmt.Println(string(table[i]))
	}
}
