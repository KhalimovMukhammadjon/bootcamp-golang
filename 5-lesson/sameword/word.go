package sameword



func RepeatWord(repetitions int, value string) string {
	var newValue string
	for i := 0; i < repetitions; i++ {
		//fmt.Print(value)
		newValue += value
	}
	return newValue
}
