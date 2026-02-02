package geometry

import "testing"

func TestCircle_Area(t *testing.T) {
	tests := []struct {
		name   string
		radius float64
		want   float64
	}{
		{"Normal int", 2.0, 12.566370614},
		{"Zero Radius", 0, 0},
		{"Normal float", 4.5, 63.617251235},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{Radius: tt.radius}
			got := c.Area()
			if !almostEqual(tt.want, got) {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
	tests := []struct {
		name   string
		radius float64
		want   float64
	}{
		{"Normal int", 5, 31.415926536},
		{"Zero Radius", 0, 0},
		{"Normal float", 3.7, 23.247785637},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{Radius: tt.radius}
			got := c.Perimeter()
			if !almostEqual(tt.want, got) {
				t.Errorf("Perimeter() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestCircle_String(t *testing.T) {
	tests := []struct {
		name   string
		radius float64
		want   string
	}{
		{"Normal int", 5, "Circle [Radius: 5.00]"},
		{"Zero Radius", 0, "Circle [Radius: 0.00]"},
		{"Normal float", 3.7, "Circle [Radius: 3.70]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{Radius: tt.radius}
			got := c.String()
			if tt.want != got {
				t.Errorf("String() = %s, want = %s", got, tt.want)
			}
		})
	}
}
