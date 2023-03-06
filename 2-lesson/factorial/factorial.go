package factorial

import "fmt"

func Factorial() {
	var a int
	count := 1
	fmt.Println("Enter number:")
	fmt.Scanln(&a)
	for i := 1; i <= a; i++ {
		count *= i
	}
	fmt.Println("Result:", count)
}
