package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// type StorageI interface {
// 	GetUser() (resp []User, err error)
// }

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Country   string `json:"country"`
}

// type UserI struct {
// 	User StorageI
// }

// func (u *UserI) GetUser() (resp []User, err error) {
// 	for _, v := range resp {
// 		fmt.Println(v)
// 	}
// 	return resp, err
// }

func main() {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var user []User
	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, v := range user {
		fmt.Println(v.Id, v.FirstName, v.LastName, v.Age, v.Country)
	}
}
