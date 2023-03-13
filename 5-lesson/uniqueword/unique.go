package uniqueword

import "fmt"

func Unique() {
	word := "users"
	arr := []int{}

	for i := 0; i < len(word); i++ {
		count := 0
		for k := 0; k < len(word); k++ {
			if word[i] == word[k] {
				count++
			}
		}
		arr = append(arr, count)
	}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			fmt.Println(string(word[i]))
			return
		}
	}
}
