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

	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()

		if got != expected {
			t.Errorf("expected %.2f but got %.2f", expected, got)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		expected := 72.0
		checkArea(t, rectangle, expected)

	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793
		checkArea(t, circle, expected)
	})
}
