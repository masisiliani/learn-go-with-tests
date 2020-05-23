package main

import (
	"fmt"
	"reflect"
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

func TestSumAll(t *testing.T) {
	t.Run("sucess to sum all items from slices and return new slice with results", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 5})
		expected := []int{3, 8}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}

func ExampleSum() {
	fibonacciNumbers := []int{1, 2, 3, 5, 8, 13}
	output := Sum(fibonacciNumbers)
	fmt.Println(output)
	// Output: 32
}
