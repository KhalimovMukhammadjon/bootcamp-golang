package main

import (
	"Muhammadjon/bootcamp/5-lesson/findnumber"
	"Muhammadjon/bootcamp/5-lesson/sameword"
	"Muhammadjon/bootcamp/5-lesson/stringvalue"
	"Muhammadjon/bootcamp/5-lesson/trimming"
	"Muhammadjon/bootcamp/5-lesson/uniqueword"
	"Muhammadjon/bootcamp/5-lesson/word"

	"fmt"
)

func main() {
	//camelcase.CamelCase()
	k := sameword.RepeatWord(5, "Hello")
	fmt.Println(k)

	t:=word.SameWord("abcde","cdegh")
	fmt.Println(t)

	stringvalue.Value()

	a := "there are lots of"
	kl := len(a)-1
	fmt.Println(kl)
	trimming.Trimming()
	uniqueword.Unique()
	findnumber.FindNumber()
}
