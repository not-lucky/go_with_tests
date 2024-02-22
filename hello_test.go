package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("test")
	want := "Hello, test"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestNewHello(t *testing.T) {
	t.Run("greet people", func(t *testing.T) {
		got := Hello("lucky")
		want := "Hello, lucky"

		assertCorrectMessage(t, got, want)
	})

	t.Run("greet no name", func(t *testing.T) {
		got := Hello("")
		want := "Hello, worlda"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
