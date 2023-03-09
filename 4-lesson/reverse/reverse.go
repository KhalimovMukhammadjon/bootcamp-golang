package reverse

import "fmt"

func ReverseNum() {
	var s int = 3456
	result := 0

	for s > 0 {
		res := s % 10
		//fmt.Println("REVERSE", res)
		result = (result * 10) + res
		s /= 10
	}
	fmt.Println(result)
}

// number 12345 // You need find index 1-2-4 ; continue with index
