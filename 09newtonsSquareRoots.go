// Author: Matthew Shiel
// Date: 01-10-2017
// Go problem 08: Merge List and Sort
// Adapted from 9. on https://data-representation.github.io/problems/go-fundamentals.html

package main

import (
	"fmt"
)

func squareNewt(z float64, x float64) float64{
	return z - ((z*z - x) / (2 * z))
}

func main() {
	x := 133.4 // Number to find the square root of
	z := 2.0 // Initial guess

	// Run code while z doesn't equal the current guess
	for z = 2.0; z != squareNewt(z, x); z = squareNewt(z, x) {
		fmt.Printf("The current guess is: %2f", z) // print current guess
	}

}