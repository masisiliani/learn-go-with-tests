package interation

import "strings"

// Repeat function return a char repeated 5 times
func Repeat(char string, qttRepeat int) string {

	return strings.Repeat(char, qttRepeat)

	// deprecated code.
	// var output string
	// for i := 0; i < qttRepeat; i++ {
	// 	output += char
	// }

	// return output
}
