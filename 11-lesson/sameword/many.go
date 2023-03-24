package sameword

import "fmt"

func FindWord1() {
	str := "password write with code w wow" //wow wow  w2 o1
	//target := "wow"
	var countW, countO int

	target := "wow"
	minW := 2
	minO := 1

	count := 0

	var newArr []string

	for i := 0; i < len(str); i++ {

		if string(str[i]) == "w" {
			countW++
			newArr = append(newArr, "w")
		}

		if string(str[i]) == "o" {
			countO++
			newArr = append(newArr, "o")
		}
	}

	for i := 0; i < len(newArr); i++ {
		if minW == 2 && minO == 1 {
			//fmt.Println("Count",count)
			count++
		}
	}

	fmt.Println("count W:", countW)
	fmt.Println("count O:", countO)

	fmt.Println(newArr)
	fmt.Println("Result", count/len(target))
	//fmt.Println("Arr",countW)
}

func FindWord() {
	str := "wo rd of write o of off wo" //woof
	var countW, countO, countF int

	target := "woof"
	minW := 2
	minO := 1
	minF := 1

	count := 0

	var newArr []string

	for i := 0; i < len(str); i++ {

		if string(str[i]) == "w" {
			countW++
			newArr = append(newArr, "w")
		}

		if string(str[i]) == "o" {
			countO++
			newArr = append(newArr, "o")
		}

		if string(str[i]) == "f" {
			countF++
			newArr = append(newArr, "f")
		}
	}

	for i := 0; i < len(newArr); i++ {
		if minW == 2 && minO == 1 && minF == 1 {
			count++
		}
	}

	fmt.Println("count W:", countW)
	fmt.Println("count O:", countO)
	fmt.Println("count F:", countF)

	fmt.Println(newArr)
	fmt.Println("Result", count/len(target))
}
