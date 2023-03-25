package discriminant

import (
	"fmt"
	"math"
)

func FindDisc() {
	var a float64 = 2
	var b float64 = 5
	var c float64 = 3

	discriminant := math.Pow(b, 2) - (4 * a * c)

	if discriminant < 0 {
		fmt.Println("No real roots")
	} else if discriminant == 0 {
		root := (-b) / (2 * a)
		fmt.Printf("One real root: %f", root)
	} else {
		root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		fmt.Printf("Two real roots: %f and %f", root1, root2)
	}
}

func Discriminant() {

	var a, b, c float64 // 2 5 3
	d := (b * b) - (4 * a * c)

	fmt.Println("Enter first number:")
	fmt.Scanln(&a)

	fmt.Println("Enter second number:")
	fmt.Scanln(&b)

	fmt.Println("Enter third number:")
	fmt.Scanln(&c)

	if d > 0 {
		fmt.Println("Roots are real and different.")
		root1 := (-b + math.Sqrt(d)) / (2 * a)
		root2 := (-b - math.Sqrt(d)) / (2 * a)
		fmt.Println("The roots are:")
		fmt.Println("First Root", root1)
		fmt.Println("Second Root", root2)
	} else if d == 0 {
		fmt.Println("Roots are equal and same.")
		root1 := (-b + math.Sqrt(d)) / (2 * a)
		fmt.Println("The root is", root1)
	} else {
		fmt.Println("Roots are complex.")
		realPart := -b / (2 * a)
		imaginaryPart := math.Sqrt(math.Abs(d)) / (2 * a)
		fmt.Println("The roots are:")
		fmt.Println("First Root", realPart, "+", "i", imaginaryPart)
		fmt.Println("Second Root", realPart, "-", "i", imaginaryPart)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in ", r)
		}
	}()

}
