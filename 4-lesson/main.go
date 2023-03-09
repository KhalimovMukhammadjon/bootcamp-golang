package main

import (
	"Muhammadjon/bootcamp/4-lesson/bignumber"
	"Muhammadjon/bootcamp/4-lesson/prime"
	"Muhammadjon/bootcamp/4-lesson/reverse"
	"fmt"
)

//evennumber "Muhammadjon/bootcamp/4-lesson/evenNumber"

func main() {
	//exercise.Number()
	//evennumber.Sum()
	reverse.ReverseNum()
	bignumber.FindBigNumber()
	v := prime.PrimeNum(15)
	fmt.Println(v)
}
