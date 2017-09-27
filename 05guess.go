// Author: Matthew Shiel
// Date: 26-09-2017
// Go problem 05: Guessing Game

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(int64(time.Now().Nanosecond())) // Call Seed using nanoseconds for a different random int every execution
	var randNum = ((rand.Int() % 100) + 1)    // Limit random number from 0 to 100

	guess := -1     // Start guess at negative to initialise While
	numOfTries := 0 // Try count
	prevGuess := 0  // Last guess the user made

	for guess != randNum {
		fmt.Print("Guess the number between 0 and 100: ") // Read guess
		fmt.Scanf("%d", &guess)

		numOfTries++ // Increase try count

		if guess == prevGuess {
			numOfTries-- // Consecutive identical tries ignored
		}

		if guess == randNum {
			fmt.Printf("Congratulations you guessed right! The number is %d \n", randNum)
		} else if guess < randNum {
			fmt.Println("Nope sorry that guess was too low!")
		} else if guess > randNum {
			fmt.Println("Nope sorry that guess was too high!")
		} else {
			fmt.Println("Oops That's not a number!")
		}

		prevGuess = guess
	}

	fmt.Printf("Wow it took you %d tries!\n", numOfTries)
}
