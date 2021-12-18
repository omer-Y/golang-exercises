/*
	pass_fail reports whether a grade is passing or failing.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() { // The "main" function gets invoked when the program launches.
	fmt.Print("Enter a grade: ")          // Prompt the user to enter a percentage grade.
	reader := bufio.NewReader(os.Stdin)   // Create a bufio.Reader, which lets us read keyboard input.
	input, err := reader.ReadString('\n') // Read what the user types, up until they press Enter.
	if err != nil {                       // If there is an error, print the message and exit.
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)            // Trim the new line character off the input
	grade, err := strconv.ParseFloat(input, 64) // Convert the input string to a float64 (numeric) value
	if err != nil {                             // If there is an error, print the message and exit.
		log.Fatal(err)
	}

	var status string // Declare the "status" variable here, so it's in scope for the rest of the function
	if grade >= 60 {  // If the grade is 60 or over, set the status to "passing". Otherwise, set it to "failing".
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status) // Print the entered grade and the pass / fail status
}
