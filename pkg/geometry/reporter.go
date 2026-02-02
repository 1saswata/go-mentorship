package geometry

type Reporter interface {
	Report(shapeName string, area float64)
}
