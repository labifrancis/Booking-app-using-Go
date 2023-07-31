package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// Declare bookings as an empty slice instead of fixed array
	var bookings []string

	for {
		// Declare variables
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

		if userTickets <= remainingTickets {
			//remaoning tickets
			remainingTickets = remainingTickets - userTickets

			// Append user's name to bookings slice
			bookings = append(bookings, firstName+" "+lasttName)

			// Thank you message
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lasttName, userTickets, email)

			//Tickets remaining
			fmt.Printf("%v tickets remaing for %v \n", remainingTickets, conferenceName)

			// Slice to get first names
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				//var firstName = names[0]
				firstNames = append(firstNames, names[0])
			}
			// All our bookings
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {

				//end program
				println("Our conference is booked full. Come back next year!")
				break

			}

		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)

		}

	}

}
