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
	var randNum = ((rand.Int() % 10) + 1)     // Limit random number from 0 to 10

	guess := -1 // Start guess at negative to initialise While
	count := 0

	for guess != randNum {
		fmt.Print("Enter your guess: ")
		fmt.Scanf("%d", &guess)

		if guess == randNum {
			fmt.Printf("Wow congratulations you guessed right! The number is %d \n", randNum)
		} else if guess < randNum {
			fmt.Println("Nope sorry that guess was too low!")
		} else if guess > randNum {
			fmt.Println("Nope sorry that guess was too high!")
		} else {
			fmt.Println("Oops try again!")
		}
	}
}
