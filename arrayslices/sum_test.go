package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("sucess to sum five numbers for array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 5, 8}
		got := Sum(numbers)
		expected := 19

		if got != expected {
			t.Errorf("expected %d but got %d, %v", expected, got, numbers)
		}
	})

	t.Run("sucess to sum tree numbers with slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("expected %d but got %d, %v", expected, got, numbers)
		}
	})
}
