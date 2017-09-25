// Author: Matthew Shiel
// Date: 25-09-2017
// Go problem 04: Factorial Digit Sum

package main

import (
        "fmt"      // Imports the IO package
        "math/big" // Imports the package necessary to declare larger data types
        "strconv"  // Package for converting digits to strings
)

func facDigitSum(x int) int {
        fact := big.NewInt(1) // Declares a new int64 to hold the larger value

        for i := 1; i <= x; i++ {
                fact.Mul(fact, big.NewInt(int64(i))) // Calculates factorial of user input by looping mutliplication 'x' times
        }

        result := 0 // Holds result of digit sum
        digit := 0 // Holds the digit 

        for i := range fact.String() {
                digit, _ = strconv.Atoi(string(fact.String()[i])) // Convert our string/character to an int
                result += digit // Calculate result by adding our digits
        }
        return result
}

func main() {
        var a int = 100 // Value we want the factorial digit sum of
        fmt.Println(facDigitSum(a)) // Calls function and prints answer
}