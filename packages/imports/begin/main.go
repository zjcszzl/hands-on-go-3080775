// packages/imports/begin/main.go
package main

// import the fmt package from the standard library
import (
	"fmt"
	"time"
)

// import the time package from the standard library

func main() {
	// use the fmt package to print the string "Hello Gopher!"
	fmt.Println("Hello Gopher!")
	// use the time package to print the current weekday
	fmt.Printf("Today is %s\n", time.Now().Weekday())
}
