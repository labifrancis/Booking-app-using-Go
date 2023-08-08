package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

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

		isValidName := len(firstName) >= 2 && len(lasttName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			//remaoning tickets
			remainingTickets = remainingTickets - userTickets

			// Append user's name to bookings slice
			bookings = append(bookings, firstName+" "+lasttName)

			// Thank you message
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lasttName, userTickets, email)

			//Tickets remaining
			fmt.Printf("%v tickets remaing for %v \n", remainingTickets, conferenceName)

			// call function to print first names
			printFirstNames(bookings)

			if remainingTickets == 0 {

				//end program
				println("Our conference is booked full. Come back next year!")
				break

			}

		} else {

			if !isValidName {
				fmt.Println("Invalid first name or last name. Please enter your name with at least 2 characters.")
			}
			if !isValidEmail {
				fmt.Println("Invalid email. Please make sure it contains the @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Printf("Invalid ticket number. You requested %v tickets but only %v tickets remaining.\n", userTickets, remainingTickets)
			}

		}

	}

}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames(bookings []string) {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		//var firstName = names[0]
		firstNames = append(firstNames, names[0])
	}
	// All our bookings
	fmt.Printf("The first names of bookings are: %v\n", firstNames)
}
