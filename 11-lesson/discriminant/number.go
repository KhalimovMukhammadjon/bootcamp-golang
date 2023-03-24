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