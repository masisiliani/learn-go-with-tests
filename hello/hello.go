package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const english = "English"
const french = "French"

// Hello function
func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	switch language {
	case english:
		return englishHelloPrefix + name
	case spanish:
		return spanishHelloPrefix + name
	case french:
		return frenchHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}

func main() {
	fmt.Println(Hello("Marcela", ""))
}
