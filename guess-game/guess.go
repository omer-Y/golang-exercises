package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seconds := time.Now().Unix() // Get the current date and time, as an integer
	rand.Seed(seconds)           // Seed the random number generator
	// Call rand.Intn to generate a random number
	target := rand.Intn(100) + 1 // Add 1 to make it an intefer from 1 to 100.
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can zou guess it?")
	fmt.Println(target)
}
