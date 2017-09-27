// Author: Matthew Shiel
// Date: 26-09-2017
// Go problem 05: Guessing Game

package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {

	rand.Seed(int64(time.Now().Nanosecond())) // Call Seed using nanoseconds for a different random int every execution
	var randNum = ((rand.Int() % 2) + 1) // Limit random number to 1000

	guess := -1 // Start guess at negative to initialise While

	for guess != randNum {
		fmt.Print("Enter your guess: ")
		fmt.Scanf("%d", &guess)

		if guess == randNum {
			fmt.Printf("Wow congratulations you guessed right! The number is %d \n", randNum)
		} else {
			fmt.Println("Oops try again!")
		}

	}

}
