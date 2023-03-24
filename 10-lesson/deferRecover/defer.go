package deferrecover

import "fmt"

func Word() {
	s := 0
	fmt.Scanln(&s)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Word", r)
		}
	}()

	if s < 0 {
		panic(s)
	}

	arr := []int{}
	arr = append(arr, s)
	fmt.Println("Arr", arr)
}

func Second(s int) {
	if s > 4 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", s))
	}

	defer fmt.Println("Defer in Second", s)
	fmt.Println("Printing in Second")
	Second(s + 1)
}
