package storage

import (
	"Muhammadjon/bootcamp/14-lesson/restaurant/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var ID, SelectFood string
var Sum, Result, Balance, OrderId int

var VarBasket []model.Basket

type Info model.Information

func (u *Info) GetUser() model.User {
	for _, v := range u.Users {
		fmt.Println(v)
	}
	return model.User{}
}

func (u *Info) GetBasket() model.Basket {
	for _, v := range u.Basket {
		fmt.Println("Basket func")
		fmt.Println("ID:", v.OrderID)
		fmt.Println("Foods:", v.Foods)
		fmt.Println("Price:", v.Sum)
		fmt.Println("Balance:", v.Balance)
		fmt.Println("")
	}
	return model.Basket{}
}

func (u *Info) GetUserByName(user model.User) model.User {
	var data model.User
	fmt.Println(user.Name)
	for _, v := range u.Users {
		if v.Name == user.Name {
			
			data = model.User{ID: v.ID, Name: v.Name, Surname: v.Surname, Balance: v.Balance}
			//data = append(data, v)
		}
		//return userss[]
	}
	return data
}

func (p *Info) GetProduct() model.Product {
	for _, v := range p.Products {
		fmt.Println("Name:", v.Name)
		fmt.Println("Price:", v.Price)
		fmt.Println("-------------------------------")
	}
	return model.Product{}
}

func (p *Info) GetProductByName(str model.Product, read string) ([]model.Product, []string) {
	var NewArr []string

	fmt.Println("read this code", str.Name)
	fmt.Println("write", read)
	fmt.Println("What would you like to eat:")

	t := 0
	for t < 3 {
		fmt.Println("What would you like to eat:")
		fmt.Scanln(&read)
		for _, v := range p.Products {
			if v.Name == read {
				//v.balance
				Result += v.Price
				Total := fmt.Sprintf("%v %v", read, Result)
				NewArr = append(NewArr, Total)
				fmt.Println("New arr", NewArr)
			}
		}

		fmt.Println("", NewArr)
		if read == "exit" {
			basket1 := model.Basket{
				OrderID: OrderId,
				Sum:     Result,
				Foods:   NewArr,
				Balance: Balance - Result,
			}
			VarBasket = append(VarBasket, basket1)
	
			break
		}
	}
	return []model.Product{}, NewArr
}

func Get() {
	data2, err := ioutil.ReadFile("../data/user.json")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	var user []model.User
	err = json.Unmarshal(data2, &user)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	data, err := ioutil.ReadFile("../data/product.json")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	var product []model.Product
	err = json.Unmarshal(data, &product)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	data3, err := ioutil.ReadFile("../data/basket.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	//var varBasket []Basket
	err = json.Unmarshal(data3, &VarBasket)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var str string
	fmt.Println("1.Order\n2.Receipt")
	fmt.Scanln(&str)

	AllProduct := &Info{
		Products: product,
	}
	// AllProduct.GetProduct()

	AllUser := &Info{
		Users: user,
	}
	//AllUser.GetUser()

	AllBasket := &Info{
		Basket: VarBasket,
	}

	if len(VarBasket) == 0 {
		OrderId = 1
	} else if len(VarBasket) != 0 {
		OrderId = len(VarBasket) + 1
	}

	switch str {
	case "1":
		userById := model.User{
			Name: "Ibrohim", // var name string // &scan
		}
		AllUser.GetUserByName(userById)

		AllProduct.GetProduct()
		selectFoodByName := model.Product{
			// Name: "Chips",
			Name: SelectFood,
		}

		_, NewArr := AllProduct.GetProductByName(selectFoodByName, SelectFood)
		AllProduct.GetProduct()

		//case "2":
		fmt.Println("Receipt")
		for _, value := range VarBasket {
			for _, valueFood := range value.Foods {
				fmt.Println(valueFood)
			}
			fmt.Println("Total:", value.Balance)
			fmt.Println()
			AllBasket.GetBasket()
		}
		basket1 := model.Basket{
			OrderID: OrderId,
			Sum:     Result,
			Foods:   NewArr,
			Balance: Balance - Result,
		}

		VarBasket = append(VarBasket, basket1)
		data3, err = json.Marshal(VarBasket)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}

		err = ioutil.WriteFile("../data/basket.json", data3, 0644)
		if err != nil {
			fmt.Println("Error writing JSON to file:", err)
			return
		}

	}
}
