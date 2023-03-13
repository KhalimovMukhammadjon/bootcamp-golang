package ticket

import (
	"fmt"
)

type User struct {
	FirstName, LastName, Email string
	TicketCount                int
}

// func WriteUser() {
// 	var firstName, lastName, email string
// 	var ticketNum int
// 	fmt.Println("Enter your firstname")
// 	fmt.Scanln(&firstName)

// 	fmt.Println("Enter your lastname")
// 	fmt.Scanln(&lastName)

// 	fmt.Println("Enter your email")
// 	fmt.Scanln(&email)

// 	fmt.Println("Enter your ticketnum")
// 	fmt.Scanln(&ticketNum)
// 	//return firstName, lastName, email, ticketNum
// }

func Ticket() {
	fmt.Println("This conference is about programming")
	var ticketNum = 50
	var result int
	var firstName, lastName, email string
	var remains int
	//newArray := []string{}
	firstNames := make([]string, 0)

	//result = ticketNum - remains
	for ticketNum >= 0 {

		fmt.Println("Enter your firstname:")
		fmt.Scanln(&firstName)

		fmt.Println("Enter your lastname:")
		fmt.Scanln(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scanln(&email)

		fmt.Println("Enter your ticketnum:")
		fmt.Scanln(&remains)

		ticketNum -= remains
		fmt.Println(result)
		firstNames = append(firstNames, firstName)

		if ticketNum == 0 {
			break
		}
		if ticketNum < 0 {
			fmt.Println("No tickets")
		} else {
			fmt.Println("You can buy it", ticketNum)
		}
	}

	//var Users []User

	// var UserData = make(map[string]string)
	// UserData["Firstname"] = firstName
	// UserData["Lastname"] = lastName
	// UserData["Email"] = email
	// UserData["Phone number"] = strconv.Itoa(remains)
	// fmt.Println(UserData)
	//UserData := make([]string, 0)

	u1 := []User{
		// {
		// 	FirstName:   "Mukhammadjon",
		// 	LastName:    "Khalimov",
		// 	Email:       "test1234@gmail.com",
		// 	TicketCount: 3,
		// },
		{
			FirstName:   firstName,
			LastName:    lastName,
			Email:       email,
			TicketCount: remains,
		},
	}

	//u1 = append(u1,)
	//Users = append(Users, u1)

	// UserData = append(UserData, firstName, lastName)
	// fmt.Println(UserData,firstNames)
	fmt.Println(firstNames)

	//sendTicket(0,"","","")
}
