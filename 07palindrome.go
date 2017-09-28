// Author: Matthew Shiel
// Date: 28-09-2017
// Go problem 07: Palindrome Tester
// Adapted from https://play.golang.org/p/AbgW-o-2MK

package main

import (
	"fmt"
	"strings"
)

func isPalindrome(input string) bool {
	// Change to lower case to solve case insensitivity
	input = strings.ToLower(input)
	// Loop half the word
	for i := 0; i < len(input)/2; i++ {
		// Compare front letter and corresponding end letter of the string
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	word := ""
	fmt.Print("Enter a word to see if it's a palindrome!: ")
	fmt.Scanf("%s", &word)
	fmt.Println(isPalindrome(word)) //Check palindrome

}
