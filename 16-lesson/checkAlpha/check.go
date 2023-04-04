package checkalpha

import (
	"fmt"
)

func CheckAlpha() {
	//arr2 := make(map[int]string)
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	word := "thequickbrownfoxjumpsoverthelazydog"
	
	arr := []string{}
	for i := 0; i < len(alphabet); i++ {
		//fmt.Println(string(word[i]))
		for k := 1; k < len(word); k++ {
			if alphabet[i] == string(word[k]) {
				fmt.Println(string(word[k]))
			}
		}
	}
	fmt.Println(arr)
}

/*



 */
