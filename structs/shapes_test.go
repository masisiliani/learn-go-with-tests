package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	expected := 40.0

	if got != expected {
		t.Errorf("expected %.2f but got %.2f", expected, got)
	}
}

func TestArea(t *testing.T) {

	scenarios := []struct {
		description string
		shape       Shape
		expected    float64
	}{
		{
			description: "Rectangle scenario",
			shape:       Rectangle{Width: 12.0, Height: 6.0},
			expected:    72.0},
		{
			description: "Circle scenario",
			shape:       Circle{Radius: 11},
			expected:    314.1592653589793},
		{
			description: "Triangle scenario",
			shape:       Triangle{Base: 12, PHeight: 6},
			expected:    36.0,
		},
	}

	for _, scenario := range scenarios {

		t.Run(scenario.description, func(t *testing.T) {
			got := scenario.shape.Area()

			if got != scenario.expected {
				t.Errorf("%#v expected %.2f but got %.2f", scenario.shape, scenario.expected, got)
			}
		})
	}
}

// func TestArea(t *testing.T) {

// 	checkArea := func(t *testing.T, shape Shape, expected float64) {
// 		t.Helper()
// 		got := shape.Area()

// 		if got != expected {
// 			t.Errorf("expected %.2f but got %.2f", expected, got)
// 		}
// 	}

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12.0, 6.0}
// 		expected := 72.0
// 		checkArea(t, rectangle, expected)

// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		expected := 314.1592653589793
// 		checkArea(t, circle, expected)
// 	})
// }
