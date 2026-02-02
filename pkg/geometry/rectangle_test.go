package geometry

import (
	"testing"
)

func TestRectangle_Area(t *testing.T) {
	tests := []struct {
		name   string
		width  float64
		height float64
		want   float64
	}{
		{"Normal 2x5", 2.0, 5.0, 10.0},
		{"Zero Width", 0, 2.0, 0},
		{"Normal 8x3", 8.0, 3.0, 24.0},
		{"Zero Height", 4.0, 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := Rectangle{Width: test.width, Height: test.height}
			got := r.Area()
			if !almostEqual(got, test.want) {
				t.Errorf("Area() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	tests := []struct {
		name   string
		width  float64
		height float64
		want   float64
	}{
		{"Normal 2x5", 2.0, 5.0, 14.0},
		{"Zero Width", 0, 2.0, 4},
		{"Normal 8x3", 8.0, 3.0, 22.0},
		{"Zero Height", 4.0, 0, 8.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := Rectangle{Width: test.width, Height: test.height}
			got := r.Perimeter()
			if !almostEqual(got, test.want) {
				t.Errorf("Perimeter() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestRectangle_String(t *testing.T) {
	tests := []struct {
		name   string
		width  float64
		height float64
		want   string
	}{
		{"Normal 2x5", 2.0, 5.0, "Rectangle [Width: 2.00, Height: 5.00]"},
		{"Zero Width", 0, 2.0, "Rectangle [Width: 0.00, Height: 2.00]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{Height: tt.height, Width: tt.width}
			if tt.want != r.String() {
				t.Errorf("String() = %s, want = %s", r, tt.want)
			}
		})
	}
}
