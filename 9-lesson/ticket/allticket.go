package ticket

import (
	"Muhammadjon/bootcamp/9-lesson/helper"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	//Age int // if else
	Email          string
	AmountOfTicket int
	PhoneNumber    string
}

var firstname, lastname, email, phonenumber string
var amountOfTicket int
var ticketNum = 50

func WriteUser(firstname, lastname, email, phonenumber string) (string, string, string, string) {
	fmt.Println("Enter your firstname:")
	fmt.Scanln(&firstname)

	fmt.Println("Enter your lastname:")
	fmt.Scanln(&lastname)

	fmt.Println("Enter your email:")
	fmt.Scanln(&email)

	fmt.Println("Enter your ticketnum:")
	fmt.Scanln(&amountOfTicket)
	return firstname, lastname, email, phonenumber
}

func CreateUser() {
	var People []Person
	// var firstname, lastname, email string
	// var amountOfTicket int
	// var ticketNum = 50
	//info := make(map[string]string)

	isValidName, isValidEmail, isValidTicketNum, isValidPhoneNum := helper.ValidateUserInput(firstname, lastname, email, phonenumber, amountOfTicket, ticketNum)
	if isValidName && isValidEmail && isValidTicketNum && isValidPhoneNum {
		// WriteUser()
		if ticketNum == 0 {
			fmt.Println("Our conference is booked out")
		}
	} else {
		if !isValidName {
			fmt.Println("firstname or lastname is too short")
		}
		if !isValidEmail {
			fmt.Println("you should enter email with @")
		}
		if !isValidTicketNum {
			fmt.Println("Number of tickets is valid")
		}
		if !isValidPhoneNum {
			fmt.Println("You should enter first 998")
		}
	}
	fmt.Println("Enter details for person")
	for ticketNum >= 0 {
		//t := WriteUser()

		ticketNum -= amountOfTicket
		person := Person{
			FirstName:      firstname,
			LastName:       lastname,
			Email:          email,
			AmountOfTicket: amountOfTicket,
		}
		if ticketNum == 0 {
			fmt.Println("User", People)
			People = append(People, person)
			break
		} else if ticketNum < 0 {
			fmt.Println("No tickets")
		} else {
			fmt.Println("You can buy it", ticketNum)
		}
		// for i := range key {
		// 	info[key[i]] = value[i]
		// }
		People = append(People, person)
	}

	for i := range People {
		fmt.Println("i", i)
	}
	fmt.Println("First", People)
}
