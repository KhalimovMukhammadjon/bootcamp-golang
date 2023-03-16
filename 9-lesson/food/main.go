package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Food struct {
	ID    int
	Name  string
	Price int
}

type NewFood struct {
	ID    int
	Sum   int
	Foods []string
}

var newArr []string

func Order() {

	var menu, selectFood string
	var selectNum int
	// var result bool
	// var foodRes []string

	fmt.Println("1.Menu\n2.Receipt\n3.Exit")
	fmt.Scanln(&menu)

	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the JSON data into a slice of Person structs
	var food []Food
	err = json.Unmarshal(data, &food)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	switch menu {
	case "1":
		for _, value := range food {
			fmt.Println(value.Name, value.Price)
		}
		for true {
			fmt.Println("What would you like to eat:")
			fmt.Scanln(&selectFood)

			fmt.Println("How much do you want to buy:")
			fmt.Scanln(&selectNum)

			for i := 0; i < len(food); i++ {
				if food[i].Name == selectFood {
					result := food[i].Price * selectNum
					fmt.Println("Total", result)
					//fmt.Printf("%v %v X %v = %v", selectFood, selectNum, food[i].Price, result)
					Total := fmt.Sprintf("%v %v X %v = %v", selectFood, selectNum, food[i].Price, result)
					newArr = append(newArr, Total)
				}
			}
			if selectFood == "exit" {
				break
			}
		}
	case "2":
		fmt.Println("Receipt")
	}
}

// lavash 4 X 21 000 = 84 000
// haggi 4 X 21 000 = 84 000

// func Sum(foodRes []string, food []Food) {
// 	var sum int
// 	findSum1 := make(map[string]int)
// 	for i := 0; i < len(foodRes); i++ {
// 		for k := 0; k < len(food); k++ {
// 			if foodRes[i] == food[k].Name {
// 				sum += food[k].Price
// 			}
// 			//fmt.Println("Summ",sum)
// 		}
// 		//fmt.Printf("select f%v", selectFood)
// 	}
// 	for i := 0; i < len(findSum1); i++ {
// 		fmt.Println(findSum1)
// 	}

// 	// fmt.Printf("%s", foodRes[0])
// 	// fmt.Println(sum)
// 	fmt.Println("new arrrr", newArrr)

// }

var total []int

func main() {
	Order()
	fmt.Println("Total", newArr)
}
