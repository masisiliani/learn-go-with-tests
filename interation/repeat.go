package interation

// Repeat function return a char repeated 5 times
func Repeat(char string) string {
	var output string
	for i := 0; i < 5; i++ {
		output += char
	}

	return output
}
