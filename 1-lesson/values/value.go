package values

import (
	"fmt"
)

var Article string = "Hello World"

func NewValue(n int) {
	//fmt.Println(n)
}

func PrimeNum() {
	var num, count int
	fmt.Scanln(&num)

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			count++
		}
		// fmt.Println(count)
	}
	if count == 2 {
		fmt.Println("It is prime number")
	}
}

func Palindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		k := len(str) - 1 - i
		if str[i] == str[k] {
			return true
		}
	}
	return false
}


func Square(a int) {
	area := a * a
	fmt.Println(area)

	p := 4 * a
	fmt.Println(p)
}

func FindYear() {
	value := "2022"
	var res string = "No beautiful"
	for i := 0; i < len(value); i++ {
		for v := i + 1; v < len(value); v++ {
			fmt.Println(string(value[v]))
			if string(value[i]) == string(value[v]) {
				res = "Beautiful"
			}
		}
	}
	fmt.Println(res)
}

func FindBigNumber() {
	num := 12345
	//num1 := str
	var count int
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(count, num)
	}
}

func BigNumber() {
	num := 12344
	var n1 int
	arr := []int{}
	// arr := num[0]
	// for i := 0; i < len(num); i++ {
	// 	fmt.Println(num,arr)
	// }
	for num > 0 {
		n1 = num % 10
		num /= 10
		arr = append(arr, n1)
	}
	fmt.Println(n1)
	//fmt.Println(res)
	//arr = append(arr,n1,n2)
	fmt.Println(arr)

	//	for s > 0 {
	// 		res = s % 10
	// 		s /= 10
	// 		count += res
	// 		fmt.Println("RES", res)
	// 	}
}

// func BigNumber() {
// 	num := "12344"
// 	var n1, n2 string
// 	arr := []int{}

// 	if len(num)%2 == 0 {
// 		n1 = string(num[0])
// 		n2 = string(num[len(num)-1])
// 	}
// 	fmt.Println(n1, n2, arr)
// }
