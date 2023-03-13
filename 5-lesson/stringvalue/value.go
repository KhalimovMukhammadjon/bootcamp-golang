package stringvalue

import "fmt"

func Value() {
	var count = -1
	var word string = "There are"
	k := "r"
	for i := 0; i < len(word); i++ {
		if string(word[i]) == k {
			count++
			if count > 0 {
				fmt.Println(i)
				return
			}
		}
	}
}
