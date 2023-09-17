package main

import "fmt"

func main() {
	// converting from a string literal
	t1 := []byte("ABCDE")

	// by a byte slice literal
	t2 := []byte{'A', 'B', 'C', 'D', 'E'}
	t3 := []byte{65, 66, 67, 68, 69}

	// by copy function
	var t4 = make([]byte, 5)
	copy(t4, "ABCDE")

	fmt.Println(t1) // [65 66 67 68 69]
	fmt.Println(t2) // [65 66 67 68 69]
	fmt.Println(t3) // [65 66 67 68 69]
	fmt.Println(t4) // [65 66 67 68 69]
}
