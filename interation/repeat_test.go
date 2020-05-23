package interation

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("success to repeat char 'a' for 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		expected := "aaaaa"

		if got != expected {
			t.Errorf("the function Repeat() doesn't work. expected %q but got %q", expected, got)
		}
	})

	t.Run("success to repeat char 'b' for 7 times", func(t *testing.T) {
		got := Repeat("b", 7)
		expected := "bbbbbbb"

		if got != expected {
			t.Errorf("the function Repeat() doesn't work. expected %q but got %q", expected, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	example := Repeat("w", 3)
	fmt.Println(example)
	// Output: www
}
