package unique

import "fmt"

// func UniqueNum() {
// 	num := make(map[int]int)
// 	num1 := []int{1, 2, 3, 4, 5, 6, 3, 8}

// 	for k,v := range num1 {
// 		num[v] = k
// 		fmt.Println("----------------------")

// 		s,err := num[v]
// 		fmt.Println(s,err)
// 		// if _,ok := num[v]; ok{
// 		// 	fmt.Println(v)
// 		// }
// 	}

// 	// Second Version
// 	// m := make(map[int]int)
// 	// m[0] = 1
// 	// m[1] = 8
// 	// m[2] = 4
// 	// m[3] = 8
// 	// num := make(map[int]int)

// 	// for k, v := range m {
// 	// 	val, err := num[v]
// 	// 	if err != true {
// 	// 		num[v] = k
// 	// 	} else {
// 	// 		fmt.Println(k, m[k], " ", val, m[val])
// 	// 	}
// 	// }
// }

// func Unique() {
// 	arr := []int{5, 6, 9, 5, 3, 6}
// 	newMap := make(map[int][]int)

// 	for k, v := range arr {
// 		newMap[v] = append(newMap[v], k)

// 		for s := range arr {
// 			if len(newMap[k]) > 0 {
// 				continue
// 			}
// 			if k == s {
// 				newMap[k] = append(newMap[k], s)
// 			} else {
// 				continue
// 			}
// 		}
// 	}
// 	for key, value := range newMap {
// 		if len(value) == 1 {
// 			delete(newMap, key)
// 		} else {
// 			continue
// 		}
// 	}
// 	fmt.Println(newMap)
// }

func Unique() {
	arr := []int{2, 4, 5, 2, 4}
	m := make(map[int][]int)
	//newArr := []int{}

	for i := 0; i < len(arr); i++ {
		newArr := []int{}
		for k := i + 1; k < len(arr); k++ {
			if arr[i] == arr[k]{
				newArr = append(newArr, i,k)
				m[arr[i]] = newArr
			}
		}
	}

	// repeat
	// for k, v := range arr {
	// 	newArr := []int{}
	// 	for s, t := range arr {
	// 		s++
	// 		if arr[k] == t{
	// 			newArr = append(newArr, k,s)
	// 			m[v] = newArr
	// 		}
	// 	}
	// }
	fmt.Println(m)
}
