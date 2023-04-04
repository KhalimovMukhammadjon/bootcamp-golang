package example

import "fmt"

func CommonFactors(a int, b int) int {
	count := 0
	//1 <= a, b <= 1000
	for i := 1; i < a; i++ {
		if a%i == 0 && b%i == 0 {
			count++
			fmt.Println("A", count)
		} else if b%i == 0 && a%i == 0 {
			count++
			fmt.Println("B", count)
		}
	}
	// for i := 1; i < b; i++ {

	// }
	return count
}
