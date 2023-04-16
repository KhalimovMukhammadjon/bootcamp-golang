package example

import "fmt"

func Find(a int) bool {
	fmt.Println(a)

	for a > 2 { // && a%8 != 0
		if a%4 != 0 {
			return false
		} else if a%8 == 0 {
			return false
		}
		a /= 4
	}
	return true
}
