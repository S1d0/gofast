package helper

// Exported function
// By making name start with capital letter, make it available for all packages in the app
func ValidateUserInput(name string, userTickets uint, remainingTickets uint) (bool, bool) {
	var isValidName bool = len(name) >= 2
	var isValidTicketNumber bool = userTickets >= remainingTickets
	return isValidName, isValidTicketNumber
}
