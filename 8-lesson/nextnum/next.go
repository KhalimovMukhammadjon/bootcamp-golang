package nextnum

import "fmt"

func FindNextNum() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	i1, i2, i3 := 0, 1, 2

	for true {
		i1++
		i2++
		i3++
		if i1 == len(arr){
			i1 = 0
		}
		if i2 == len(arr){
			i2 = 0
		}
		if i3 == len(arr){
			i3 = 0
		}
		fmt.Println(arr[i1], arr[i2], arr[i3])
	}
}
