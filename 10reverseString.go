// Author: Matthew Shiel
// Date: 01-10-2017
// Go problem 10: Reverse String
// Adapted from https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-
// Adapted from https://rosettacode.org/wiki/Reverse_a_string#Go

package main

import (
	"fmt" // Import IO package
)

func main() {
	s := "I love data representation and querying" // Word to reverse

	fmt.Printf("The reverse of the String '%s' is '%s'", s, reverse(s)) // Print string in reverse
}

func reverse(s string) string {
	// input should be ASCII UTF-8 encoded string
	// create an array of runes
	r := make([]rune, len(s))
	l := len(s)
	// Loop and add runes to r array
    for _, rune := range s {
		l--
		r[l] = rune
	}
	// Return reversed string
    return string(r)
}
