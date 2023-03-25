package shape

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float32
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius
}

// func (r Rectangle) () float64 {
// 	return r.Width * float64(r.Height)
// }

func (r Rectangle) perimeter() float64 {
	return r.Width * float64(r.Height)
}

func PrintShapeInfo(s Shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}
