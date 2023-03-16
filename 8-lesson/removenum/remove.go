package removenum

import (
	"fmt"
	"reflect"
	"strings"
)

func RemoveNum() {
	arr := "38012732"
	s := []string{}
	str := "3"

	for i := 0; i < len(arr); i++ {
		if string(arr[i]) == str {
			fmt.Println("Delete")
			continue
		} else {
			s = append(s, string(arr[i]))
			fmt.Println("Append")
		}
	str2 := strings.Join(s, " ")
	fmt.Println(str2)
	fmt.Println(reflect.TypeOf(str2))
	}
	fmt.Println(arr)
	//fmt.Println(str2)
}
