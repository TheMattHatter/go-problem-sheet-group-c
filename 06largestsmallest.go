// Author: Matthew Shiel
// Date: 28-09-2017
// Go problem 06: GLargest and Smallest Int

package main

import (
	"fmt"
)

func main() {
	list := [10]int{4, 9, 50, 1, 234, 56, 78, 34, 20, 89} // List of integers
	var smallest = list[0]
	var largest = list[0]

	for i := range list { 
			if list[i] < smallest {
				smallest = list[i]
			}

			if list[i] > largest {
				largest = list[i]
			}
	}

	fmt.Printf("The smallest number in the list is %d\n", smallest)
	fmt.Printf("The largest number in the list is %d\n", largest)
	
}