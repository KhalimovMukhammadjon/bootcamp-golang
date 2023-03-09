package main

import (
	"Muhammadjon/bootcamp/5-lesson/sameword"
	"Muhammadjon/bootcamp/5-lesson/stringvalue"
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
}
