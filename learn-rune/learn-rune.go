package main

import "fmt"

func main() {

	t := []rune("hello world")
	fmt.Println(string(t))

	// update
	t[0] = 'H'
	t[6] = 'W'
	fmt.Println(string(t)) // Hello world

	// add
	t = append(t, '!', '?')
	fmt.Println(string(t)) // Hello World!?

	// delete
	t = t[:len(t)-1]
	fmt.Println(string(t)) // Hello World!
}
