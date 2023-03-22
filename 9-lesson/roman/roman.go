package roman

import "fmt"

func ChangeRoman() {
	str := "IV"
	romanMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	n := len(str) // 3

	result := romanMap[str[n-1]] // showed us the last number
	fmt.Println("result", result)

	// for i := n - 2; i >= 0; i-- {
	// 	fmt.Println("Index", i)
	// 	if romanMap[str[i]] < romanMap[str[i+1]] {
	// 		result -= romanMap[str[i]]
	// 		fmt.Println("Truee",result)
	// 	} else {
	// 		result += romanMap[str[i]]
	// 		fmt.Println("new", romanMap[str[i]])
	// 		fmt.Println("tr", result)
	// 	}
	// }
	//fmt.Println(result)

	for i := 0; i < len(str)-1; i++ {
		fmt.Println("Index", len(str)-1)

		fmt.Println("----------------------------------------------")

		fmt.Println("First", romanMap[str[i]])
		fmt.Println("Second", romanMap[str[i+1]])

		if romanMap[str[i]] < romanMap[str[i+1]] {
			result -= romanMap[str[i]]
			fmt.Println("if:", result)
		} else {
			result += romanMap[str[i]]
			fmt.Println("else:", result)
		}
	}
	fmt.Println(result)
}

func RomanToInt() {
	str := "VII"
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	s := len(str)
	result := romanMap[str[s-1]] // romanMap[str[len(str)-1]] // 
	fmt.Println("last num:",result)

	for i := 0; i < len(str)-1; i++ {
		fmt.Println("index",i)

		fmt.Println("----------------------------------------------")
		fmt.Println(romanMap[str[i]])
		fmt.Println("TEST",romanMap[str[i+1]])

		if romanMap[str[i]] < romanMap[str[i+1]] {
			fmt.Println("r", result)
			result -= romanMap[str[i]]
			fmt.Println("r", result)
		} else {
			result += romanMap[str[i]]
			fmt.Println("else", result)
		}
	}
	fmt.Println("sum", result)
}
