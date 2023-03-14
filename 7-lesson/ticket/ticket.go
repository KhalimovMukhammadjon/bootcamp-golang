package ticket

import (
	"fmt"
)

type User struct {
	FirstName   string
	LastName    string
	Email       string
	TicketCount int
}

var Bookings = make([]User, 0)

func NewTicket() {
	fmt.Println("This conference is about programming")
	var ticketNum = 50
	var result int
	var firstName, lastName, email string
	var remains int

	firstNames := make([]string, 0)
	allNames := make([]string, 0)

	//var UserData = make(map[string]string)
	// UserData["Lastname"] = lastName
	// UserData["Email"] = email
	// UserData["Phone number"] = strconv.Itoa(remains)
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
		//v := GetFirstNames()
		Bookings = append(Bookings, )

		allNames = append(allNames, firstName, lastName, email)
		fmt.Println("All names", allNames)

		if ticketNum == 0 {
			break
		}
		if ticketNum < 0 {
			fmt.Println("No tickets")
		} else {
			fmt.Println("You can buy it", ticketNum)
		}
	}
	GetFirstNames()
	fmt.Println("Name", firstNames)
	fmt.Println("All names", allNames)
}

func GetFirstNames() []string {
	firstName := []string{}
	for _, booking := range Bookings {
		firstName = append(firstName, booking.FirstName)
		//fmt.Printf("%v.%v\n",k,booking.FirstName)
		fmt.Println(firstName)
	}
	return firstName
}
