// types/conversion/begin/main.go
package main

import "fmt"

func main() {
	// declare float and convert it to an unsigned int
	a := 42.7
	b := uint(a)
	fmt.Printf("a is %v with type %T, and b is %v with type %T\n", a, a, b, b)

}
