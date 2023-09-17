package main

import "fmt"

func main() {
	/* User specified types */
	const a int32 = 12        // 32-bit integer also called rune
	const b float32 = 20.5    // 32-bit floaT
	var c complex128 = 1 + 41 // 128-bit complex number
	var d uint16 = 255        // 16-bit unsigned integer
	var e byte = 251          // same as uint8 0-255
	var f int8 = -128         // -128 to +128
	/* Default types */
	n := 42              // int
	pi := 3.14           // float64
	x, y := true, false  // bool
	z := "Go is awesome" // string

	fmt.Printf("user-specified types:\n %T %T %T %T %T %08b\n", a, b, c, d, e, f)
	fmt.Printf("default types:\n %T %T %T %T %T\n", n, pi, x, y, z)
}
