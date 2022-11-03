package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("Get your tickets here to attend")

	conferenceTickets = 30
	fmt.Println(conferenceTickets)
	
}