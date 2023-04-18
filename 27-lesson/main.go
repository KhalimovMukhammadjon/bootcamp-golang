package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var s = make(chan string)

func text(ch chan string) string {
	fmt.Println("In process")
	time.Sleep(time.Second * 2)

	// for {
	// 	str := "Hello World"
	// 	s <- str
	// 	close(s)
	// }
	str := "Hello World"
	s <- str
	close(s)
	return str
}

func read() {
	//time.Sleep(time.Second * 1)
	go text(s)
}

func main() {
	wg.Add(1)
	go read()
	for r := range s {
		fmt.Println("Result:", r)
	}
}
