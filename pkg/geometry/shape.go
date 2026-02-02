package geometry

type Shape interface {
	Area() float64
	Perimeter() float64
	String() string
}
