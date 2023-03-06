package main

import "fmt"

// func Condition() {
// 	var a, b, c int
// 	fmt.Scanln(&a)
// 	fmt.Scanln(&b)
// 	fmt.Scanln(&c)

// 	if a > b && a > c {
// 		fmt.Println("A is big value", a)
// 	} else if b > a && b > c {
// 		fmt.Println("B is big value ", b)
// 	} else if c > a && c > b {
// 		fmt.Println("C is big value ", c)
// 	} else if a == b && a == c && b == c {
// 		fmt.Println("Neither is big value ")
// 	} else {
// 		fmt.Println("null")
// 	}
// }

func FindNum() {
	var a int
	fmt.Scanln(&a)

	if a%2 == 0 {
		fmt.Println("Even number")
	} else if a%2 != 0 {
		fmt.Println("Odd number")
		if a%2 != 0 {
			fmt.Println("Will be divided", a)
		} else {
			fmt.Println("Will not divided")
		}
	}
}

func FindTotal() {
	count := 0
	var n int
	fmt.Println("Enter number")
	fmt.Scanln(&n)
	for i := 0; i <= n; i++ {
		count += i
	}
	fmt.Println(count)
}

func FindEvenTotal() {
	var a, count int
	fmt.Println("Enter number")
	fmt.Scanln(&a)
	for i := 0; i <= a; i++ {
		if i%2 == 0 {
			count += i
			//fmt.Println("Even number", count)
		}
	}
	fmt.Println("Even number:", count)
}


