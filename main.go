package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "go conference"

const conferenceTickets int = 50 //cannot change the value inside it
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	numberOfTicets uint
}

var wg = sync.WaitGroup{}

func main() {

	//in int type the number can be negetive
	greetUsers()

	//cntrl +c to exit the looping

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		//noTicketsRemsining := remainingTickets == 0 //== for compering 2 values
		if remainingTickets == 0 {
			//end program
			fmt.Printf("Our conference is booked out.Come back next year.")
			//break
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
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n ", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining tickets\n ", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	// _ is consider as blank identifier
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames //return func ca return the data as a result(can take input and output)
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	// create a map for user

	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		numberOfTicets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf(" %v remining tickets is %v\n ", conferenceName, remainingTickets)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%vtickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("###############")
	wg.Done()
}
