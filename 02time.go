// Author: Matthew Shiel
// Date: 19-09-2017
// Go problem 02: Current time
// Adapted from https://tour.golang.org/welcome/4

package main

import (
	"fmt"  // Importing the IO package fmt
	"time" // Importing the time package for for measuring and displaying time.
)

func main() {
	fmt.Println("Welcome to the playground!") // Prints our phrase

	fmt.Println("The time is", time.Now()) // Call time.Now() to fetch the current date/time
}
