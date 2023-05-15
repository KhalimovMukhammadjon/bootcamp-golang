package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello World!")
}

func main() {
	go hello()
	time.Sleep(1 * time.Second)
	goodbye()
}

func goodbye() {
	fmt.Println("Good Bye!")
}


// func main() {
//     // Create a waitgroup with 2 goroutines.
//     var wg sync.WaitGroup
//     wg.Add(2)

//     // Start two goroutines.
//     go func() {
//         defer wg.Done()
//         fmt.Println("Hello from goroutine 1!")
//     }()

//     go func() {
//         defer wg.Done()
//         fmt.Println("Hello from goroutine 2!")
//     }()

//     // Wait for the goroutines to finish.
//     wg.Wait()

//     // Print "All goroutines finished."
//     fmt.Println("All goroutines finished.")
// }
