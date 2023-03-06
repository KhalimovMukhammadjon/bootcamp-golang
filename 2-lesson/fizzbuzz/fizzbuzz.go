package fizzbuzz

import "fmt"

func FizzBuzz() {
	var n int
	fmt.Println("Enter number:")
	fmt.Scanln(&n)

	if n%3 == 0 && n%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if n%3 == 0 {
		fmt.Println("Fizz")
	} else if n%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println("Enter positive number")
	}
}
