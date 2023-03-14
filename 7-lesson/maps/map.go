package maps

import "fmt"

func CreateMap() {
	//make(map[string]string)
	var input string
	var str1, str2 string
	authors := make(map[string]string)

	authors["Alisher Navoiy"] = "Xamsa"
	authors["Abdulla Qodiriy"] = "Utkan kunlar"
	authors["James"] = "Atomic Habits"
	authors["Abdulla Qahhor"] = "Anor"

	for k, v := range authors {
		fmt.Printf("Author:%v Books:%v\n", k, v)

		fmt.Println("Enter menu:")
		fmt.Scanln(&input)

		if input == "search" {
			fmt.Println("----------------------------")
			fmt.Println("Enter author you want to search")
			fmt.Scanln(&str2)
			fmt.Println("This book`s name is", authors[str2]) // or authors["James"]
		} else if input == "delete" {
			fmt.Println("----------------------------")
			fmt.Println("Enter author you want to remove")
			fmt.Scanln(&str1)
			delete(authors, str1)
		} else if input == "exit" {
			break
		}
	}

	// second version
	// fmt.Println("----------------------------")
	// fmt.Println("Enter author you want to search")
	// fmt.Scanln(&str2)
	// fmt.Println("This book`s name is", authors[str2]) // or authors["James"]

	// fmt.Println("----------------------------")
	// fmt.Println("Enter author you want to remove")
	// fmt.Scanln(&str1)
	// delete(authors, str1)
	fmt.Println(authors)
}
