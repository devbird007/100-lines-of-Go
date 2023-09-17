package main

import (
	"fmt"
)

type Student struct {
	FirstName  string
	LastName   string
	Age        int
	School     string
	Department string
	StudentID  int
}

func main() {
	// Declaring a variable of type Student
	var John Student
	fmt.Println(John)

	// Declaring and initializing a variable using struct literal
	// We declare each field with its Key and Value
	var Mary = Student{
		FirstName:  "Mary",
		LastName:   "Raymond",
		School:     "MIT",
		StudentID:  10034,
		Age:        19,
		Department: "Computer Science",
	}
	fmt.Println(Mary)

	var Timothy = Student{
		"Timothy",
		"Jackson",
		21,
		"MIT",
		"Computer Science",
		10100,
	}
	fmt.Println(Timothy)
}
