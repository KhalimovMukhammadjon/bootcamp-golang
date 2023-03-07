package multiply

import "fmt"

func Multiply() {
	var s int
	result := 0
	fmt.Println("Enter number:")
	fmt.Scanln(&s)
	if s <= 15 {
		for i := 1; i <= s; i++ {
			fmt.Printf("--------------------\n")
			for k := 1; k <= i; k++ {
				result = i * k
				fmt.Printf("%d X %d = %d\n", i, k, result)
			}
		}
	} else {
		fmt.Println("Enter a number below 15")
	}
}
