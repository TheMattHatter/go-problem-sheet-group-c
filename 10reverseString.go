// Author: Matthew Shiel
// Date: 01-10-2017
// Go problem 10: Reverse String
// Adapted from https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go

package main

import (
	"fmt" // Import IO package
)

func main() {
	s := "I love data representation and querying" // Word to reverse

	fmt.Printf("The reverse of the String '%s' is '%s'", s, reverse(s)) // Print string in reverse
}

func reverse(s string) string {
	// Assume input is  an ASCII UTF-8 encoded string
	// create an array of runes and loop for the length of the string/array
	runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {

        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
