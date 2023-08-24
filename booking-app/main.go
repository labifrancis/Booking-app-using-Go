package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"

var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	// Declare bookings as an empty slice instead of fixed array

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTickets(userTickets, firstName, lastName, email)

			// call function to print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

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

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		//var firstName = names[0]
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	// Declare variables
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	// remaoning tickets
	remainingTickets = remainingTickets - userTickets

	// Append user's name to bookings slice
	bookings = append(bookings, firstName+" "+lastName)

	// Thank you message
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	// Tickets remaining
	fmt.Printf("%v tickets remaing for %v \n", remainingTickets, conferenceName)
}
