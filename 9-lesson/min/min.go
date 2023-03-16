package min

import "fmt"

func MinNum(arr []int) int {
	var min = arr[0]

	for k, v := range arr {
		//fmt.Println(k, v)
		if v < min {
			min = arr[k]
			//fmt.Println("MIN",min)
		}
	}
	fmt.Println("MIN", min)

	return min
}
