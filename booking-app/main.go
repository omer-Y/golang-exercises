package main

import "fmt"

func main() {
	conferenceName := "Go Conference" // another method for define variable.
	const conferenceTickets int = 50  // this variable type cannot be changed of value
	var remainingTickets uint = 50    // except negative value

	// %T return the variable type
	fmt.Printf("conferenceTickets is %T, remainingTicets is %T, conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName)

	// %v return the variable value
	fmt.Printf("Welcome to %v our conference booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	userName = "Omer"
	userTickets = 15

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
