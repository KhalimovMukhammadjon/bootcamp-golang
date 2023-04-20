package example

import "fmt"

func Find(a int) bool {
	fmt.Println(a)
	for a > 3 {
		if a%4 != 0 {
			return false
		}
		a /= 4
	}
	fmt.Println(a)
	return a == 1
}
