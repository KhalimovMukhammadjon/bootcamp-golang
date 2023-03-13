package trimming

import "fmt"

func Trimming() {
	text := "Creating kata is fun" // Creating ka...
	k := 14
	newText := "..."
	var count string

	for i := 0; i < k-3; i++ {
		//fmt.Print(string(text[i]))
		count += string(text[i])
	}
	fmt.Println(count + newText)
}
