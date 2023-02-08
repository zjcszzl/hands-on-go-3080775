// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val string = "global"

func main() {
	// create a local variable "val" and assign it an int value
	val := 3

	// print the value of the local variable "val"
	fmt.Printf("Local variable val is %v\n", val)

	// print the value of the package-level variable "val"
	print()

	// update the package-level variable "val" and print it again
	update()
	print()

	// print the pointer address of the local variable "val"
	var ptr *int = &val
	fmt.Printf("Address of pointer is %v\n", ptr)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*ptr = 12
	fmt.Printf("Local variable val is %v\n", val)
}

func print() {
	fmt.Printf("Global variable val is %v\n", val)
}

func update() {
	val = "new global"
}
