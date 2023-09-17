package main

import (
	"fmt"
)

func main() {
	var aString string = "A string"
	var strPtr = &aString

	fmt.Printf("The address of \"aString\" is %v\n", &aString)
	fmt.Printf("The value stored in \"strPtr\" is %v\n", strPtr)
	fmt.Printf("The data type of \"strPtr\" is %T\n", strPtr)

	var aFloat float64 = 9.00
	floatPtr := &aFloat

	fmt.Printf("The address of \"aFloat\" is: %v\n", &aFloat)
	fmt.Printf("The value stored in \"floatPtr\" is: %v\n", floatPtr)
	fmt.Printf("The data type of \"floatPtr\" is: %T\n", floatPtr)

}
