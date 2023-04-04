package leetcode

import "fmt"

func FindSum() {
	arr := []int{4, 6, 8}
	fmt.Println(arr)
	count := 0

	arr1 := []int{}

	for i := range arr {
		//v = i + 1
		count = arr[i] + i + 1
		if count > 9 {
			s := count % 10
			arr1 = append(arr1, s)

		} else {
			arr1 = append(arr1, count)
		}
		fmt.Println(count)
		//arr1 = append(arr1, count)
	}
	fmt.Println(arr1)
	//fmt.Println(arr)
	//fmt.Println(count)
}
