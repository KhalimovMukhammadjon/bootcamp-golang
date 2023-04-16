package sort

import "fmt"

func Sort() {
	a := []int{1, 2, 4}
	b := []int{1, 5, 6}
	//arr := []int{}

	//fmt.Println("ARR", a)

	// for i, v := range a {
	// 	fmt.Println(v, i)
	// 	for _, t := range b {
	// 		if v > t {
	// 			arr = append(arr, v)
	// 		}
	// 	}
	// }
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
		for k := 0; k < len(b); k++ {
			fmt.Println(b[k])
		}
	}
	//a = append(a, b...)
}

/*
// 1.Write a program that creates two Go routines, one to print even numbers from 1 to 10 and the other to print odd numbers from 1 to 10.

// 2.Write a program that creates a Go routine to read integers from standard input and send them to a channel. Create another Go routine that reads from the channel and prints the sum of the integers received.

// 3.Write a program that creates a Go routine to generate random numbers between 1 and 100 and send them to a channel. Create another Go routine that reads from the channel and prints the numbers in descending order.


/// 5.Write a program that creates a Go routine to calculate the factorial of a number using recursion. Use a channel to communicate the result of the calculation to the main program.

*/