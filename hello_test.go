package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Anish")
	want := "Hello, Anish!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
