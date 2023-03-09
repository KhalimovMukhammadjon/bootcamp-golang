package camelcase

import (
	"fmt"
	"strings"
)

func CamelCase() {
	s1 := "the-stealth-warrior"
	arr := []string{}

	for i := 0; i < len(s1); i++ {
		//fmt.Println(s1)
		if string(s1[i]) == "-" {
			i++
			arr = append(arr, strings.ToUpper(string(s1[i])))
		} else {
			arr = append(arr, string(s1[i]))

		}
	}
	fmt.Println(arr)
}
