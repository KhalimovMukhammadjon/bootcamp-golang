package squaresum

import "fmt"

func FindSum() {
	arr := []int{1, 2, 2}
	count := 0
	res := 0
	for _, v := range arr {
		count = v * v
		res += count
		// res += v * v // second version
		fmt.Println(count, v)
	}
	fmt.Println("Result", res)

}
