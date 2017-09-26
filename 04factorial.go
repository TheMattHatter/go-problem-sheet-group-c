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

        if x == 0 {
                return -1 // If the user enters 0
        } else {
                fact := big.NewInt(1) // int64 to hold the large values

                for i := 1; i <= x; i++ {
                        fact.Mul(fact, big.NewInt(int64(i))) // Calculates factorial of user input by looping mutliplication 'x' times
                }

                result := 0 // result of digit sum
                digit := 0  // individual digit

                for i := range fact.String() {
                        digit, _ = strconv.Atoi(string(fact.String()[i])) // Convert our string/character to an int
                        result += digit                                   // Sum up digit total
                }
                return result
        }
}

func main() {
        var a int = 100             
        fmt.Println(facDigitSum(a))
}