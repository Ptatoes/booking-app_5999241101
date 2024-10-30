package main

import (
	"fmt"
	"strings"
)

var conferenceName = "go conference"

const conferenceTickets = 50 //cannot change the value inside it
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	//in int type the number can be negetive
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	//cntrl +c to exit the looping
	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remainingTickets, userTickets, &bookings, firstName, lastName, email, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			//noTicketsRemsining := remainingTickets == 0 //== for compering 2 values
			if remainingTickets == 0 {
				//end program
				fmt.Printf("Our conference is booked out.Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println(" first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println(" The email is not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid")
			}

			//fmt.Printf("Your input data is invalid, try agaian")
			//continue //skip remainder of its body and restesting all the condition
		}

	}
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n ", confName)
	fmt.Printf("We have total of %v tickets and %v remaining tickets\n ", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	// _ is consider as blank identifier
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames //return func ca return the data as a result(can take input and output)
}
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user for their name
	fmt.Println("Please enter the first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter the last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter the email:")
	fmt.Scan(&email)

	fmt.Println("Please enter the number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings *[]string, firstName string, lastName string, email string, conferenceName string) {

	remainingTickets = remainingTickets - userTickets
	*bookings = append(*bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf(" %v remining tickets is %v\n ", conferenceName, remainingTickets)

}
