package main

import "testing"

func TestSubtest(t *testing.T) {
	t.Run("saying hello to the people", func(t *testing.T){
		got:= Hello("Mahesh", "Hindi")
		want:= "Namaste, Mahesh!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}