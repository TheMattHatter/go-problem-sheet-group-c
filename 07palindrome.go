// Author: Matthew Shiel
// Date: 28-09-2017
// Go problem 07: Palindrome Tester
// Adapted from https://play.golang.org/p/AbgW-o-2MK and http://www.golangpro.com/2017/04/

package main

import (
	"fmt"     // IO package
	"strings" // Strings package to manipulate strings
	"unicode" // Unicode package to check for Unicode characters
)

func isPalindrome(input string) bool {
	input = strings.ToLower(input) // Change to lower case to solve case insensitivity

	var letters []rune // Array to hold runes

	// We don't care about the index so we toss it with the blank identifier
	for _, r := range input {
		if unicode.IsLetter(r) {
			letters = append(letters)
		}
	}

	// Loop for half the word
	for i := range letters {
		// Compare front letter and corresponding end letter of the string
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	word := ""
	fmt.Print("Enter a word to see if it's a palindrome!: ")
	fmt.Scanf("%s", &word)
	fmt.Println(isPalindrome(word)) // Check palindrome
}
