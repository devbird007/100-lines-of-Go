package main

import "fmt"

func main() {

	s := "中文"
	r := []rune(s)
	fmt.Println(r) // result: 20013 25991]

	p := []rune{20013, 25991}
	q := string(p)
	fmt.Println(q) // 中文
}
