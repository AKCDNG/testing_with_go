package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Anish", "")
		want := "Hello, Anish!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, World!' when an empty string is inputted", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Output in Spanish", func(t *testing.T) {
		got := Hello("Anish", "Spanish")
		want := "Hola, Anish!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Output in French", func(t *testing.T) {
		got := Hello("Anish", "French")
		want := "Bonjour, Anish!"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
