// Author: Matthew Shiel
// Date: 21-09-2017
// Go problem 03: FizzBuzz test
// Adapted from ttp://wiki.c2.com/?FizzBuzzTest

package main

import "fmt"

func main() {
	var fizz string = "Fizz"
	var buzz string = "Buzz"
	var fizzBuzz string = "FizzBuzz"

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(fizz)
		} else if i%5 == 0 {
			fmt.Println(buzz)
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println(fizzBuzz)
		} else {
			fmt.Println(i)
		}
	}
}
