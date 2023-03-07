package multiplication

import "fmt"

func Multiplication() {
	var s int
	result := 0
	fmt.Println("Enter number:")
	fmt.Scanln(&s)
	if s <= 15 {
		for i := 1; i <= 10; i++ {
			result = s * i
			fmt.Printf("%d X %d = %d\n", s, i, result)
		}
	} else {
		fmt.Println("Enter a number below 15")
	}
}
