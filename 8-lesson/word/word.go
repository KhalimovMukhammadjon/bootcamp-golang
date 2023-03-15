package word

import "fmt"

func FindWord() {
	var count string

	var s int
	fmt.Scanln(&s)

	// arr := []int{}
	arr := make([]string, s)

	for i := 0; i < s; i++ {
		fmt.Scanln(&arr[i])
		count += arr[i]
	}
	fmt.Println(count,len(count))
}
