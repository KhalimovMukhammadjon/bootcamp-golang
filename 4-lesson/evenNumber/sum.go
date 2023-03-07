package evennumber

import (
	"fmt"
	"strconv"
)

func Sum() {
	// 4 5 6 7 8  == 18 - 12
	var s int = 45678
	str := strconv.Itoa(s)
	sum1 := 0
	sum2 := 0

	for i := 0; i < len(str); i++ {
		if i%2 == 0 {
			a, _ := strconv.Atoi(string(str[i]))
			sum1 += a
			fmt.Println(sum1)
		} else {
			a, _ := strconv.Atoi(string(str[i]))
			sum2 += a
			fmt.Println(sum2)
		}
	}
	fmt.Println(sum1 - sum2)
}
