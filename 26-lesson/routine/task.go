package routine

import (
	"fmt"
	"sync"
	"time"
)

/*
// func GO() {
// 	// 1.Write a programls that creates two Go routines, one to print even numbers from 1 to 10 and the other to print odd numbers from 1 to 10.

// 	// 2.Write a program that creates a Go routine to read integers from standard input and send them to a channel. Create another Go routine that reads from the channel and prints the sum of the integers received.

// 	// 3.Write a program that creates a Go routine to generate random numbers between 1 and 100 and send them to a channel. Create another Go routine that reads from the channel and prints the numbers in descending order.

// 	// / 5.Write a program that creates a Go routine to calculate the factorial of a number using recursion. Use a channel to communicate the result of the calculation to the main program.
// }
*/

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for j := range jobs {
// 		fmt.Printf("Worker %d started job %d\n", id, j)
// 		time.Sleep(time.Second)
// 		fmt.Printf("Worker %d finished job %d\n", id, j)
// 		results <- j * 2
// 	}
// }

// func main() {
// 	jobs := make(chan int, 100)
// 	results := make(chan int, 100)

// 	for w := 1; w <= 3; w++ {
// 		go worker(w, jobs, results)
// 	}

// 	for j := 1; j <= 9; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)

// 	for a := 1; a <= 9; a++ {
// 		<-results
// 	}
// }
var wg sync.WaitGroup

func EvenNum() {
	fmt.Printf("Go routine is running\n")
	time.Sleep(time.Second * 1)

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("This is even number:", i)
		}
	}
	wg.Done()
}

func OddNum() {
	fmt.Printf("Go routine is running\n")
	time.Sleep(time.Second * 1)

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println("This is odd number:", i)
		}
	}
	wg.Done()
}
