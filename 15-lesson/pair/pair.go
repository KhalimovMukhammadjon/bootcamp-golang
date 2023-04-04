package pair

import "fmt"

func FindPair() {
	total := make(map[string]int)
	arr := []string{"red", "red", "red", "green"}
	for _, v := range arr {
		fmt.Println(v)
		var countA int
		// if v == "red" {
		// 	countR++
		// }
		// if v == "blue" {
		// 	countB++
		// }
		// res = countB + countR
		for _, k := range arr {
			if v == k {
				countA++
				total[string(v)] = countA
			}
		}

	}
	result := 0
	for _, t := range total {
		if t%2 == 0 {
			//fmt.Println("RRRR")
			result += t
		} else if t%2 != 0 && t > 1 {
			result += t-1
		}
	}
	fmt.Println("Result", result/2)
	fmt.Println("Total", total)
	// fmt.Println(countR, countB, res)
	// fmt.Println(res / 2)
}
