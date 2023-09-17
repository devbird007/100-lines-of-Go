package main

import (
	"fmt"
	"reflect"
)

func main() {
	/* Define an array containing programming languages */
	languages := [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust", // Must contain trailing comma
	}

	/* Define slices */
	classics := languages[0:3]  // alternatively languages[:3]
	modern := make([]string, 4) // len(modern) = 4
	modern = languages[3:7]     // include 3 exclude 7
	new := languages[7:9]       // alternatively languages[7:]

	fmt.Printf("classic languages: %v\n", classics)
	fmt.Printf("modern languages: %v\n", modern)
	fmt.Printf("new languages: %v\n", new)

	// -- snip -- //
	allLangs := languages[:]                     // copy of the array
	fmt.Println(reflect.TypeOf(allLangs).Kind()) // slice
	fmt.Println("%v\n", allLangs)

	/* Create a slice containing web frameworks */
	frameworks := []string{
		"React", "Vue", "Angular", "Svelte",
		"Laravel", "Django", "Flask", "Fiber",
	}

	jsFrameworks := frameworks[0:4:4]         // length 4 capacity 4
	frameworks = append(frameworks, "Meteor") // not possible with arrays

	fmt.Printf("all frameworks: %v\n", frameworks)
	fmt.Printf("js frameworks: %v\n", jsFrameworks)

}
