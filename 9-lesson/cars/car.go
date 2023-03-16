package cars

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Car struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Price int    `json:"price"`
}

func CreateCar() {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	var cars []Car

	err = json.Unmarshal(data, &cars)
	if err != nil {
		fmt.Println("Error unmarshaling", err)
		return
	}

	for _, value := range cars {
		fmt.Println(value.Name, value.Color, value.Price)
	}

	// add new car
	cars = append(cars,
		Car{
			Name:  "Bugatti",
			Color: "Blue",
			Price: 450000,
		},
		Car{
			Name:  "Aston Martin",
			Color: "Green",
			Price: 230000,
		},
	)

	// update cars
	for i, val := range cars {
		if val.Name == "Aston Martin" {
			cars[i].Color = "Black"
		}
	}

	data, err = json.Marshal(cars)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// write json data to a file
	err = ioutil.WriteFile("cars/data.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

}
