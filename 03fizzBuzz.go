// Author: Matthew Shiel
// Date: 21-09-2017
// Go problem 03: FizzBuzz test
// Adapted from ttp://wiki.c2.com/?FizzBuzzTest

package main

import "fmt" // Imports the IO package

func main() {
	var fizz string = "Fizz" // fizz string
	var buzz string = "Buzz" // buzz string

	// for loop loops 100 times
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println(fizz + buzz) // Prints FizzBuzz For numbers which are multiples of both three and five
		} else if i % 3 == 0 {
			fmt.Println(fizz) // Prints fizz for multiples of three
		} else if i % 5 == 0 {
			fmt.Println(buzz) // Prints buzz for multiples of five
		} else {
			fmt.Println(i) // Prints the current number in the loop
		}
	}
}
