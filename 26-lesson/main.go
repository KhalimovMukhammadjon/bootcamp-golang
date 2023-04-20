package main

import (
	//"Muhammadjon/bootcamp/26-lesson/routine"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func evenNum() {
	fmt.Printf("Go routine is running\n")
	time.Sleep(time.Second * 1)

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("This is even number:", i)
		}
	}
	wg.Done()
}

func oddNum() {
	fmt.Printf("Go routine is running\n")
	time.Sleep(time.Second * 1)

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println("This is odd number:", i)
		}
	}
	wg.Done()
}

func findSum() {
	ch := make(chan []int)
	receiver := make(chan int)
	res := 0

	go func() {
		n := []int{3, 4, 5, 2, 8}
		ch <- n
		close(ch)

		for i := 0; i < len(n); i++ {
			res += n[i]
			receiver <- res
		}
		//fmt.Println("result:", res)
		close(receiver)
	}()

	go func() {
		for read := range receiver {
			fmt.Println("result", read)
		}
	}()

	for val := range ch {
		fmt.Println("val", val)
	}
}

func main() {
	// wg.Add(1)
	// go evenNum()
	// time.Sleep(time.Second * 3)
	// wg.Wait()

	// wg.Add(1)
	// go oddNum()
	// time.Sleep(time.Second * 3)
	// wg.Wait()

	// fmt.Println("All Go routines are done")
	findSum()
}

/*
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	fmt.Println("nn", n*factorial(n-1))
	return n * factorial(n-1)
}

func calculateFactorial(n int, c chan int) {
	result := factorial(n)
	c <- result
}

func main() {
	n := 5
	c := make(chan int)
	go calculateFactorial(n, c)
	fmt.Printf("Factorial of %d is %d", n, <-c)
}
*/
