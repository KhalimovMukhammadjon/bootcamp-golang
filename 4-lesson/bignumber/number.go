package bignumber

import (
	"fmt"
)

func FindBigNumber() {
	var n int = 56784
	arr := []int{}
	var count int
	//newArr := []int{}

	//var result, count int
	// for n > 0 {
	// 	count = n % 10
	// 	fmt.Println(count)
	// 	result = (result * 10) + count
	// 	n /= 10
	// }
	// fmt.Println(result)

	// FIRST version
	// for i:=0; i<len(str); i++{
	// 	if i%2 != 0{
	// 		continue
	// 	}
	// 	//strconv.Atoi(string(str[i]))
	// 	//fmt.Println(str[len(str[:])-1])
	// }
	fmt.Println(n)
	for n > 0 {
		count = n % 10
		n /= 10
		arr = append(arr, count)
	}
	fmt.Println(arr)

	// for i,v := range arr{
	// 	if i%2 != 0{
	// 		continue
	// 	}
	// 	fmt.Println("Continue",i,v)
	// 	newArr = append(newArr, v)
	// 	fmt.Println(newArr)
	// }

	min := arr[0]
	max := arr[0]
	for _, v := range arr {
		//fmt.Println(i, v)
		if v < min {
			min = v
		} else  if v > max{
			max = v
		}
	}
	fmt.Println("MIN", min)
	fmt.Println("MAX", max)

}
