package main

import "testing"

func Test_Hello(t *testing.T) {

	got := Hello()
	want := "Hello, word"

	if got != want {
		t.Errorf("got %q and want %q", got, want)
	}

}
