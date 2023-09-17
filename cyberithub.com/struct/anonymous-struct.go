package main

import (
	"fmt"
)

func main() {
	// Declaring and initializing an anonymous struct
	mary := struct {
		firstName  string
		lastName   string
		age        int
		school     string
		department string
	}{
		"Mary",
		"Raymond",
		19,
		"MIT",
		"Computer Science",
	}
	fmt.Println(mary)
}
