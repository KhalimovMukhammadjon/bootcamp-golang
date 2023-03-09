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

// func SameWord(word1, word2 string) string {
// 	//"abcde","cdegh"
// 	for i := 0; i < len(word2); i++ {
// 		//v := string(word2[0])

// 		if word2[0] == word1[i] {
// 			fmt.Println(string(word1[i:]))
// 			// word2[:len(copy)]
// 			res := string(word1[i:])
// 			if res == word2[:len(res)] {
// 				return word1[:i] + word2
// 			}
// 		} else {
// 			return word1 + word2
// 		}
// 	}
// 	return word1 + word2
// }

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
