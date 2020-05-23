package main

// Sum function receive an array with 5 numbers and return the sum result
func Sum(numbers []int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}
