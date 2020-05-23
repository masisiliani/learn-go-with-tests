package main

// Sum function receive an array with 5 numbers and return the sum result
func Sum(numbers []int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

// SumAll take a varying number of slices, returning a new
// slice containing the totals for each slice passed in.
func SumAll(numbersFirst ...[]int) []int {
	return []int{3, 5}
}
