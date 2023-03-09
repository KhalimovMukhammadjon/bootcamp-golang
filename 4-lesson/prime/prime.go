package prime

import "fmt"

// func PrimeNum(s int) bool {
// 	for i := 2; i <= s; i++ {
// 		isPrime := true
// 		for k := 2; k < i; k++ {
// 			if i%k == 0 {
// 				isPrime = false
// 			}
// 		}
// 		if isPrime == true {
// 			fmt.Println(i)
// 		}
// 	}
// 	return true
// }

func PrimeNum(s int) bool {
	var count int
	isPrime := true
	for i := 2; i < s/2; i++ {
		if s%i == 0 {
			count++
		}
	}
	if count == 0 && s != 1 {
		fmt.Println(s, "is a Prime Number")
		isPrime = true
		return isPrime
	} else {
		fmt.Println(s, "is not a Prime Number")
		isPrime = false
		return isPrime
	}
	//fmt.Println(count)
	//return false
	//return isPrime
}
