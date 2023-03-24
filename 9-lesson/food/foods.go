package food

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Food struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Basket struct {
	OrderID int      `json:"id"`
	Foods   []string `json:"food"`
	Sum     int      `json:"sum"`
}

var newArr []string
var orderId, sum, result int //added result
//var foods []string

func Menu() {

	var menu, selectFood string
	var selectNum int

	// var name string
	// var id, price int

	fmt.Println("1.Menu\n2.Receipt\n3.Exit")
	fmt.Scanln(&menu)

	data, err := ioutil.ReadFile("food/data.json")
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

	// Read Order json
	data1, err := ioutil.ReadFile("food/order.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var varBasket []Basket
	err = json.Unmarshal(data1, &varBasket)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// This code adds plus 1 to id
	if len(varBasket) == 0 {
		orderId = 1
	} else if len(varBasket) != 0 {
		orderId = len(varBasket) + 1
	}

	switch menu {
	case "1":
		for _, value := range food {
			fmt.Println(value.Name, value.Price)
		}

		for true {
			fmt.Println("What would you like to eat:")
			fmt.Scanln(&selectFood)

			for i := 0; i < len(food); i++ {
				if food[i].Name == selectFood {
					fmt.Println("How much do you want to buy:")
					fmt.Scanln(&selectNum)
					result += food[i].Price * selectNum
					//fmt.Printf("%v %v X %v = %v", selectFood, selectNum, food[i].Price, result)
					//Total := fmt.Sprintf("%v %v X %v = %v", selectFood, selectNum, food[i].Price, result)
					Total := fmt.Sprintf("%v %v x %v", selectFood, selectNum, food[i].Price)

					// fmt.Println("New total t",Total)
					// fmt.Println("TTT",total)
					newArr = append(newArr, Total)
					fmt.Println(newArr, "Neww")
				}
			}
			if selectFood == "exit" {
				basket1 := Basket{
					OrderID: orderId,
					Sum:     result,
					Foods:   newArr,
					//Quantity:
				}

				varBasket = append(varBasket, basket1)
				data1, err = json.Marshal(varBasket)
				if err != nil {
					fmt.Println("Error marshaling JSON:", err)
					return
				}

				err = ioutil.WriteFile("food/order.json", data1, 0644)
				if err != nil {
					fmt.Println("Error writing JSON to file:", err)
					return
				}
				//fmt.Println("Basket22", basket2)
				// for _, value := range basket2 {
				// 	fmt.Println(value.OrderID, value.Foods, value.Sum)
				// }

				break
			}
		}
	case "2":
		fmt.Println("Receipt")
		for _, value := range varBasket {
			fmt.Println("ID:", value.OrderID) //value.Sum
			for _, valueFood := range value.Foods {
				fmt.Println(valueFood)
			}
			fmt.Println("Total:", value.Sum)
			//result = food[i].Price * selectNum
			// Total := fmt.Sprintf("%v %v X  = %v", selectFood, selectNum, result)
			// fmt.Println(Total)
			fmt.Println()
		}
	}
}

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

func Order() {
	Menu()
	//fmt.Println("Total", newArr)
}
