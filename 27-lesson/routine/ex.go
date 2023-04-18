package main

import (
	"fmt"
	"time"
)

func count(word string, c chan string) {
	for i := 1; i < 5; i++ {
		c <- word
		time.Sleep(time.Millisecond * 500)
		//fmt.Println("Count", c)
	}
	close(c)
}

func main() {
	c := make(chan string)
	go count("horse", c)

	for msg := range c {
		fmt.Println("message:", msg)
	}
}
