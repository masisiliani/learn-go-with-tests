package structs

import "math"

// Shape is an abstract interface for geometry shape
type Shape interface {
	Area() float64
}

// Rectangle is a struct with values of rectagle object
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculate the perimeter of a rectangle given a height and width
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle is a struct with values of circle object
type Circle struct {
	Radius float64
}

// Area returns the area of a rectangle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
