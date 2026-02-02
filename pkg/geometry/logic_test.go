package geometry

import "testing"

type MockReporter struct {
	calledName string
	calledArea float64
}

func (m *MockReporter) Report(s string, area float64) {
	m.calledName = s
	m.calledArea = area
}

func TestProcessShape(t *testing.T) {
	tests := []struct {
		name          string
		inputShape    Shape
		wantArea      float64
		wantShapeName string
	}{
		{
			"Rectangle 2x5",
			Rectangle{Width: 2, Height: 5},
			10.0,
			"Rectangle [Width: 2.00, Height: 5.00]",
		},
		{
			"Circle Radius 10",
			Circle{Radius: 10},
			314.159265359,
			"Circle [Radius: 10.00]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var myMock MockReporter
			ProcessShape(tt.inputShape, &myMock)
			if !almostEqual(myMock.calledArea, tt.wantArea) {
				t.Errorf("Area() got:= %2f, want:= %2f", myMock.calledArea, tt.wantArea)
			}
			if myMock.calledName != tt.wantShapeName {
				t.Errorf("Shape Name got:= %q, want %q", myMock.calledName, tt.wantShapeName)
			}
		})
	}

}
