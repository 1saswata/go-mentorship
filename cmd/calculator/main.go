package main

import (
	"fmt"

	"github.com/1saswata/go-mentorship/pkg/geometry"
)

func main() {
	shapes := []geometry.Shape{
		geometry.Rectangle{Width: 3, Height: 8},
		geometry.Circle{Radius: 5},
		geometry.Rectangle{Width: 3.7, Height: 5.5},
		geometry.Circle{Radius: 8.5},
	}
	for _, shape := range shapes {
		fmt.Printf("%s\nArea: %.2f\nPerimeter: %.2f\n\n", shape, shape.Area(), shape.Perimeter())
	}
}
