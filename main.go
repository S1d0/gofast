package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var userName string
	var userTickets uint
	var bookings []string

	fmt.Printf("Welcome to our %v booking app \n", conferenceName)
	fmt.Println("We have total of ", conferenceTickets, "and,", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attends")

	for {
		fmt.Println("Enter your first name: ")
		fmt.Scan(&userName)

		fmt.Println("How many tickets for you?: ")
		fmt.Scan(&userTickets)

		bookings = append(bookings, userName)
		remainingTickets = remainingTickets - userTickets
		fmt.Printf("User %vorder %v ticktes \n", userName, userTickets)

		// For-Each loop
		var names []string
		for _, booking := range bookings {
			names = append(names, strings.ToUpper(booking))
		}
		fmt.Printf("Remaining tickets %v\n, These are all our bookings:%v\n", remainingTickets, names)

		if remainingTickets == 0 {
			fmt.Println("Our conf is booked out")
			break
		}
	}
}
