package main

import "fmt"

func User() {
	// fmt.Println("Golang")
	// var (
	// 	v1  bool = true
	// 	v2  string = "New"
	// 	v3  int16 = 23232
	// 	v4  int32 = 535636346
	// 	v5  int64 = 4343434343434949494
	// 	v6  uint16 = 32323
	// 	v7  uint32 = 3439434394
	// 	v8  uint64 = 3595459459459050505
	// 	v9  rune = 243434343
	// 	v10 float32 = 364654.323
	// 	v11 float64 = 43563456.232
	// 	v12 byte = 42
	// 	v13 int = 2323932929393933933
	// )
	// fmt.Printf("Type: %T - Value: %v\n", v1, v1)
	// fmt.Printf("Type: %T - Value: %v\n", v2, v2)
	// fmt.Printf("Type: %T - Value: %v\n", v3, v3)
	// fmt.Printf("Type: %T - Value: %v\n", v4, v4)
	// fmt.Printf("Type: %T - Value: %v\n", v4, v5)
	// fmt.Printf("Type: %T - Value: %v\n", v6, v6)
	// fmt.Printf("Type: %T - Value: %v\n", v7, v7)
	// fmt.Printf("Type: %T - Value: %v\n", v8, v8)
	// fmt.Printf("Type: %T - Value: %v\n", v9, v9)
	// fmt.Printf("Type: %T - Value: %v\n", v10, v10)
	// fmt.Printf("Type: %T - Value: %v\n", v11, v11)
	// fmt.Printf("Type: %T - Value: %v\n", v12, v12)
	// fmt.Printf("Type: %T - Value: %v\n", v13, v13)
	var n1, n2 int
	var v string
	fmt.Println("Enter first number")
	fmt.Scanln(&n1)
	fmt.Println("Enter second number")
	fmt.Scanln(&n2)

	fmt.Println("Enter: + - * /")
	fmt.Scanln(&v)
	if v == "+" {
		fmt.Println(n1 + n2)
	} else if v == "-" {
		fmt.Println(n1 - n2)
	} else if v == "/" {
		fmt.Println(n1 / n2)
	} else if v == "*" {
		fmt.Println(n1 * n2)
	} else {
		fmt.Println("Unknown type")
	}
}

// func main(){
// 	User()
// }
