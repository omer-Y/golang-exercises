package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50 // this variable type cannot be changed of value
	var remainingTicets = 50

	fmt.Println("Welcome to", conferenceName, "our conference booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTicets, "are still available.")
	fmt.Println("Get your tickets here to attend")
}
