package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// Declare bookings as an empty slice instead of fixed array
	var bookings []string

	// Print details about slice

	var firstName string
	var lasttName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lasttName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets

	// Append user's name to bookings slice

	bookings = append(bookings, firstName+" "+lasttName)

	// Print details about slice
	fmt.Printf("Print the whole Slice: %v\n", bookings)
	fmt.Printf("Print the first value of the Slice: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice length: %v\n", len(bookings))
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lasttName, userTickets, email)
	fmt.Printf("%v tickets remaing for %v \n", remainingTickets, conferenceName)

}
