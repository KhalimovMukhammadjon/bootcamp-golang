package dividenumber

import "fmt"

func DivideNumber() {
	number := 121
	n := number

	count := 0
	result := 0

	for number > 0 {
		result = number % 10
		if n%result == 0 {
			count++
		}
		number /= 10
	}
	fmt.Println(count)

	//fmt.Println(count)
	// for i := 0; i < len(number); i++ {
	// 	fmt.Println(string(number[i]), i)
	// 	// if number[i]%121 == 0 {
	// 	// 	fmt.Println("count",string(number[i]))
	// 	// 	count++
	// 	// }

	// }
	// if number[i]%i == 0 && number[i]%i+1 == 0 {
	// }
}
