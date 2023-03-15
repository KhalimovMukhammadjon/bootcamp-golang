package ticket

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	//Age int // if else
	Email          string
	AmountOfTicket int
}

func CreateUser() {
	var People []Person
	var firstname, lastname, email string
	var amountOfTicket int
	var ticketNum = 50
	//info := make(map[string]string)
	fmt.Println("Enter details for person")
	for ticketNum >= 0 {
		// fmt.Printf("Enter firstname, lastname and email for person %d: ", i+1)
		// fmt.Scan(&firstname, &lastname, &email)

		fmt.Println("Enter your firstname:")
		fmt.Scanln(&firstname)

		fmt.Println("Enter your lastname:")
		fmt.Scanln(&lastname)

		fmt.Println("Enter your email:")
		fmt.Scanln(&email)

		fmt.Println("Enter your ticketnum:")
		fmt.Scanln(&amountOfTicket)

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
		//info[People[i]] = 
		fmt.Println("i", i)
	}
	fmt.Println("First", People)
}
