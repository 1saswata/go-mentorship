package geometry

func ProcessShape(s Shape, r Reporter) {
	area := s.Area()
	r.Report(s.String(), area)
}
