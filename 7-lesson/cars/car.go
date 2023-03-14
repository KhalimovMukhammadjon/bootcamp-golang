package cars

import "fmt"

type Car struct {
	Name  string
	Model string
	Color string
	Price int
}

// pointer
func NewCar(name string) *Car {
	c := Car{Name: name}
	c.Price = 120000
	fmt.Println(c.Price)
	FindCar()
	return &c
}

func FindCar() {
	car1 := Car{
		Name:  "BMW M8",
		Model: "Sedan",
		Color: "Black",
		Price: 168000,
	}
	fmt.Println(car1)
	fmt.Printf("Name:%v\nModel:%v\nColor:%v\nPrice:%d", car1.Name, car1.Model, car1.Color, car1.Price)
}
