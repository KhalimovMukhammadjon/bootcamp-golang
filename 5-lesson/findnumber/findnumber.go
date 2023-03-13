package findnumber

import "fmt"

func FindNumber(){
	num := "38012732"
	s := "8"
	for i:=0; i<len(num); i++{
		if string(num[i]) == s{
			fmt.Println(string(num[i]))
			fmt.Println(num)
		}
	}
}