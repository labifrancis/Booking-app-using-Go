package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	fmt.Printf("conferenceName is a %T, conferenceTickets is a %T, remainingTickets is a %T\n", conferenceName, conferenceTickets, remainingTickets)

	var userName string
	var userTickets int
	// ask user for their name
	fmt.Scan(&userName)

	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}
