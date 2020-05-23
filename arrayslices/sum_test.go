package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 5, 8}
	got := Sum(numbers)
	expected := 19

	if got != expected {
		t.Errorf("expected %d but got %d, %v", expected, got, numbers)
	}

}
