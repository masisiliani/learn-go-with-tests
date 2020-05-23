package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("sucess to sum any size of slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 5, 8}
		got := Sum(numbers)
		expected := 19

		if got != expected {
			t.Errorf("expected %d but got %d, %v", expected, got, numbers)
		}
	})
}

func ExampleSum() {
	fibonacciNumbers := []int{1, 2, 3, 5, 8, 13}
	output := Sum(fibonacciNumbers)
	fmt.Println(output)
	// Output: 32
}
