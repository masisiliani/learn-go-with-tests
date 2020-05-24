package structs

// Perimeter calculate the perimeter of a rectangle given a height and width
func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

// Area returns the area of a rectangle
func Area(width, height float64) float64 {
	return width * height
}
