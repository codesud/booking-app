package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to our" %v "booking application\n", conferenceName)
	fmt.Println("We have total of" %v "tickets and" %v "are still available.", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	
}