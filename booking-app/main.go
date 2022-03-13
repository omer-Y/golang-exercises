package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50 // this variable type cannot be changed of value
	var remainingTicets = 50

	fmt.Printf("Welcome to %v our conference booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTicets)
	fmt.Println("Get your tickets here to attend")
}
