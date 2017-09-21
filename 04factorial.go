// Author: Matthew Shiel
// Date: 21-09-2017
// Go problem 04: Factorial Digit Sum

package main

import "fmt"

func factorial(x int) int {
	//var a [10]int	
	sum := 1
	for i := 1; i <= x; i++ {
		sum = sum * i
	}
	return sum
}

func main() {
	var a int = 10
	fmt.Println(factorial(a))
}