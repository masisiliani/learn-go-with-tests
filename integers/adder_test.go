package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	expected := 4

	if got != expected {
		t.Errorf("fails to sum two numbers: expected %d but got %d", expected, got)
	}
}

// Add function should be sum two number
// For to do it if we pass 1 and 6 as argument, we will recieve 7 as return
func ExampleAdd() {
	sum := Add(1, 6)
	fmt.Println(sum)
	// Output: 7
}
