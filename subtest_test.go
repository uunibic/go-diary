package main

import "testing"

func TestSubtest(t *testing.T) {
	t.Run("saying hello to the people", func(t *testing.T){
		got:= Hello("Mahesh")
		want:= "Hello, Mahesh!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World!' if an empty string is provided", func(t *testing.T){
		got:= Hello("")
		want:= "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}