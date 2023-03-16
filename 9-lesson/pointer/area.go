package pointer

import "fmt"

type Rect struct {
	width  int
	height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func (r Rect) perimetr() int {
	return 2*r.width + 2*r.width
}

func FindArea() {
	res := Rect{
		width:  10,
		height: 20,
	}
	fmt.Println(res.area())
	fmt.Println(res.perimetr())
}
