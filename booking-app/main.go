package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	"strings"
	//"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings = make([]map[string]string, 0)

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTickets(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			// call function to print first names
			firstNames := getFirstNames(bookings)
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

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	for _, booking := range bookings {
		//var firstName = names[0]
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
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

func bookTickets(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	// remaoning tickets
	remainingTickets = remainingTickets - userTickets

	// craete map for a user \
	var userData = make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// Append user's name to bookings slice
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v")

	fmt.Printf("List of bookings is %v\n", bookings)

	// Thank you message
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	// Tickets remaining
	fmt.Printf("%v tickets remaing for %v \n", remainingTickets, conferenceName)
}
