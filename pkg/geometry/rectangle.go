package geometry

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle [Width: %.2f, Height: %.2f]", r.Width, r.Height)
}
