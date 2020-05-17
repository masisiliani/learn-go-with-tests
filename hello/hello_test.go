package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q and want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Marcela", "English")
		want := "Hello, Marcela"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World', when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)

	})

	t.Run("speak in Spanish", func(t *testing.T) {
		got := Hello("Marcela", "Spanish")
		want := "Hola, Marcela"

		assertCorrectMessage(t, got, want)
	})

	t.Run("speak in French: say 'Bonjour, Marcela', when send 'Marcela' and 'French'", func(t *testing.T) {
		got := Hello("Marcela", "French")
		want := "Bonjour, Marcela"

		assertCorrectMessage(t, got, want)
	})
}
