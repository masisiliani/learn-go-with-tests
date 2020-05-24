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
func SumAll(numbersSlices ...[]int) []int {
	var sliceSumAll []int
	for _, numbers := range numbersSlices {
		sliceSumAll = append(sliceSumAll, Sum(numbers))
	}

	return sliceSumAll
}

// SumAllTails calculates the totals of the "tails" of each slice
func SumAllTails(numbersSlices ...[]int) []int {
	var sliceSumTails []int
	for _, numbers := range numbersSlices {
		if len(numbers) == 0 {
			sliceSumTails = append(sliceSumTails, 0)
		} else {
			tail := numbers[1:]
			sliceSumTails = append(sliceSumTails, Sum(tail))
		}
	}
	return sliceSumTails
}
