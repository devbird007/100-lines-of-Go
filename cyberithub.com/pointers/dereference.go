package main

import (
	"fmt"
)

func main() {
	var aString string = "A string"
	var strPtr = &aString

	fmt.Printf("The address of \"aString\" is %v\n", &aString)
	fmt.Printf("The value stored in \"strPtr\" is %v\n", strPtr)

	// Dereferencing operation to return the value pointed to by strPtr
	fmt.Printf("The value stored in the variable \"strPtr\" points to %q\n", strPtr)

	// Indirectly update aString through strPtr
	*strPtr = "An updated string"
	fmt.Printf("The value stored in the variable \"strPtr\" points to %q\n", strPtr)
	fmt.Printf("The value of \"aString\" is %q\n", aString)
}
