package findword

func CountOfSpace(line string) int {
	i := 0
	for _, value := range line {
		if value == ' ' {
			i++
		} else {
			break
		}
	}
	return i
}
