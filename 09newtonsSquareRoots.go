// Author: Matthew Shiel
// Date: 01-10-2017
// Go problem 09: Netwon's Square Root
// Adapted from 9. on https://data-representation.github.io/problems/go-fundamentals.html

package main

import (
	"fmt"  // Import IO package
	"math" // Math package for sqrt
)

func squareNewt(z float64, x float64) float64 {
	if z == 0 {
		return 0
	}

	return z - (z*z-x)/(2*z) // Formula for Netwon's Square Root
}

func main() {
	x := 87.0 // Number to find the square root of
	z := 1.0  // Initial guess

	// Run code while z doesn't equal the current guess
	for z = 1.0; z != squareNewt(z, x); z = squareNewt(z, x) {
		fmt.Printf("The current guess is: %2.6f\n", z) // print current guess
	}

	// Print the approximation according to Netwon's Method
	fmt.Printf("\n%2f is an approximation for the square root of %2f\n", z, x)

	// Print the square root calculated by math.Sqrt()
	fmt.Printf("The square root according to math.Sqrt is %2f\n", math.Sqrt(x))
}
