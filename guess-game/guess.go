package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

	reader := bufio.NewReader(os.Stdin) // Create a bufio.Reader, which lets us read keyboard input.

	fmt.Print("Make a guess: ")           // Ask for a number.
	input, err := reader.ReadString('\n') // Read what the user types, up until they press Enter.
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)  // Remove the newline.
	guess, err := strconv.Atoi(input) // Convert the input string to an integer.
	if err != nil {
		log.Fatal(err)
	}

	if guess < target {
		fmt.Println("Oops. Your guess was LOW.")
	} else if guess > target {
		fmt.Println("Oops. Your guess was HIGH.")
	}
}
