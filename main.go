package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
	"time"
)

// package level variable
const conferenceTickets uint = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

// Custom types
// lightweight version of class which doesn't support inheritance
type UserData struct {
	firstName       string
	lastName        string
	numberOfTickets uint
}

func main() {
	var userName string
	var userTickets uint

	greetUsers()

	for {
		fmt.Println("Enter your first name: ")
		fmt.Scan(&userName)

		fmt.Println("How many tickets for you?: ")
		fmt.Scan(&userTickets)

		isValidName, isValidTicketNumber := helper.ValidateUserInput(userName, userTickets, remainingTickets)

		if isValidName && isValidTicketNumber {
			bookTicket(userTickets, userName, "Nowak")
			go sendTicket(userTickets, userName, "jj@mymail.com")
			
			fmt.Printf("User %vorder %v ticktes \n", userName, userTickets)

			printFirstNames()
			fmt.Printf("Remaining tickets %v\n, These are all our bookings\n", remainingTickets)

		}
		if remainingTickets == 0 {
			fmt.Println("Our conf is booked out")
			break
		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %s\n", conferenceName)
	fmt.Println("We have total of ", conferenceTickets, "and,", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attends")
}

func printFirstNames() {
	var firstNames []string
	for _, booking := range bookings {
		var name string = strings.ToUpper(booking.firstName)
		firstNames = append(firstNames, name)
	}
	fmt.Printf("The first names: %s\n", firstNames)

}

func bookTicket(userTicket uint, firstName string, lastName string) {
	remainingTickets = remainingTickets - userTicket

	// create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		numberOfTickets: userTicket,
	}
	bookings = append(bookings, userData)
}

// Simulate slow execution
func sendTicket(userTickets uint, firstName string, email string) {
	// Stops thread for 10 seconds
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintln("%v tickets for %v", userTickets, firstName)
	fmt.Println("###############")
	fmt.Printf("Sending tickets %v to email address %v \n", ticket, email)
	fmt.Println("###############")
}
