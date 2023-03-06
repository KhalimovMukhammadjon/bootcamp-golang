package lesson2

import "fmt"

func OrderById() {
	fmt.Println("Empty Function")
}

func FindPrimeNum() {
	prime := []int{}
	for i := 0; i < 100; i++ {
		var fake = false
		for k := 2; k <= i; k++ {
			if i%k == 0 && k != i {
				fake = true
			}
		}
		if fake == false {
			prime = append(prime, i)
		}
	}

}
