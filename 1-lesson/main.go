package main

import (
	// "1-lesson/lesson2"
	"1-lesson/values"
	"fmt"
)

func main() {
	
	fmt.Println(values.Article)
	values.NewValue(4)

	// var s uint8 = 255
	// fmt.Print(s)
	// fmt.Print(s)

	// var k float64 = 9
	// y := 5.0
	// fmt.Println(math.Pow(k, y))

	// 2-Lesson
	// var s1, s2 int = 4, 5
	// var n int

	// num1,num2 := int(6),int(7)
	// num1,num2 = num2,num1

	// n = s1
	// fmt.Println(n)
	// n = s2
	// fmt.Println(n)

	// POW
	// s := 3
	// count := 1
	// for i := 1; i <= 3; i++ {
	// 	count *= s
	// }
	// fmt.Println(count)

	// var n1 int
	// fmt.Scanf("%v", &n1)
	// fmt.Printf("%d %T", n1, n1)

	// LOOP
	// var s int
	// var str, res string

	// fmt.Scanln(&s)

	// for i := 0; i < s; i++ {
	// 	fmt.Println("Enter your word")
	// 	fmt.Scanln(&str)
	// 	res += str + " "
	// }
	// fmt.Println(res)

	// PRIME NUM
	// var num, count int
	// fmt.Scanln(&num)

	// for i := 1; i <= num; i++ {
	// 	if num%i == 0 {
	// 		count++
	// 	}
	// }

	// if count == 2 {
	// 	fmt.Println("It is prime number")
	// }
	//values.PrimeNum()

	// p :=values.Palindrome("olma")
	// fmt.Println(p)

	// values.Square(8)
	//values.FindYear()
	values.BigNumber()
	values.FindBigNumber()
}
