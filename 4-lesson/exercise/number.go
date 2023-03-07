package exercise

import (
	"fmt"
	"log"
	"strconv"
)

func Number() {

	fmt.Println("Enter number")

	val := true
	count := 0
	s := "stop"

	for val {
		fmt.Scanln(&s)
		if s == "stop" {
			break
		} else {
			res, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
				return
			}
			count += res
		}

	}
	fmt.Println("Result", count)
}
