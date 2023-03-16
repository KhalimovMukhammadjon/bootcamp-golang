package helper

import "strings"

func ValidateUserInput(firstname, lastname, email, phonenumber string, amountOfTicket int, ticketNum int) (bool, bool, bool, bool) {
	isValidName := len(firstname) >= 3 && len(lastname) >= 3
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNum := amountOfTicket > 0 && amountOfTicket < ticketNum
	isValidPhoneNum := strings.Contains(phonenumber, "998")

	// 998 123 43 54
	for i := 0; i < len(phonenumber); i++ {
		if string(phonenumber[i]) == "998"{
			// phonenumber[0] == "+" && phonenumber
			continue
		}	
	}

	return isValidName, isValidEmail, isValidTicketNum, isValidPhoneNum
}
