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
		got := SumAll([]int{1, 2}, []int{3, 5, 5})
		expected := []int{3, 13}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("success to calculate the totals of the 'tails' of each slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 5, 5})
		expected := []int{2, 10}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
	t.Run("success when it pass a empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 5, 5})
		expected := []int{0, 10}

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
