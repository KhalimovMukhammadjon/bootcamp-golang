package word

func Word(s1, s2 string) {
	count := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[count] {
			count++
		} else {
			count = 0
			if s1[i] == s2[count] {
				count++
			}
		}
	}
}

//"abcde","cdegh"
func SameWord(a, b string) string {
	for i := 0; i < len(a)-1; i++ {
		if a[i] == b[0] {
			str := a[i:]
			if str == b[:len(str)] {
				return a[:i] + b
			}
		}
	}
	return a + b
}
