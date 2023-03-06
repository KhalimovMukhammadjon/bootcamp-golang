package fibonacci

import "fmt"

func Fibonnacci() {
	var n int
	fmt.Println("Enter number")
	fmt.Scanln(&n)

	n1 := 0
	n2 := 1
	next := 0
	for i := 1; i < n1; i++ {
		if i == 1 {
			continue
		}
		if i == 2 {
			continue
		}
		next = n1 + n2
		n2 = next
		fmt.Println(next)
	}
}
