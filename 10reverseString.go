// Author: Matthew Shiel
// Date: 01-10-2017
// Go problem 10: Reverse String

package main

import (
	"fmt" // Import IO package
)

func reverse(s string) string {
	reverse := ""
	// Loop through the string backwards and store in a string
    for i := len(s)-1; i >= 0; i-- {
        reverse += string(s[i])
    }
    return reverse 
}

func main() {
	s := "I love data querying and representation" // Word to reverse

	fmt.Printf("The reverse of the String '%s' is %s", s, reverse(s)) // Print string in reverse
}