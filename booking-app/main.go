package main

import (
	"fmt"
	"strings"
)

func main() {
	// var conferenceName string = "Go Conference"
	conferenceName := "Go Conference" // only applies to var and not const
	const conferenceTicket int = 50
	var remainingTickets uint = 50 // unsigned int and it will not have negative number
	// var bookings = [50]string{"Deep", "Sumi"}
	// var bookings [50]string
	var bookings []string // use slice

	// fmt.Printf("conferenceTicket is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTicket, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your ticket here to attend!")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName) // used for input , & is a pointer is a var that points to the memory address of another var

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)
		// firstName = "Deep"
		// userTickets = 2

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// isInvalidCity := city != "Singapore" || city !=" London"

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You'll receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceTicket)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			// var noTicketsRemaining bool = remainingTickets == 0/// or can be written as
			// noTicketsRemaining := remainingTickets == 0

			if remainingTickets == 0 {
				// end the program
				fmt.Printf("Our conference is booked out. Come back next year.")
				break
			}
			// }
			//  else if userTickets == remainingTickets {
			// do something else
		} else {
			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			// break
			// continue
			if !isValidName {
				fmt.Println((" first name or last name you entered is too short"))
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			// fmt.Println("Your input data is invalid, try again.")
		}
	}

	city := "London"
	switch city {
	case "New York":
		// execute code for booking new york conference tickets
	case "Singapore":
		// execute
	// case "London":
	// execute
	case "London", "Berlin": // if same logic needs to be executed
		// execute
	case "Mexico City":
		// execute
	case "Hong Kong":
		// execute
	default:
		fmt.Print("No valid city selected!")
	}
}
