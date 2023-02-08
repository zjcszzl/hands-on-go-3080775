// types/slices/begin/main.go
package main

import "fmt"

func main() {
	// declare a slice string the make function
	names := make([]string, 0)

	// append 3 names to the slice
	names = append(names, "John")
	names = append(names, "Mary")
	names = append(names, "jane")

	// print the slice
	fmt.Println(names)
	// slice the slice using syntax slice[low:high]
	fmt.Println(names[1:3])
	// [Jane Mary]
	fmt.Println(names[:])
	// [Jane Mary]
	// [John]
	// [John Jane Mary]
}
