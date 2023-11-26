package main

import "testing"

func TestDisct(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	assertStrings := func (t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got error %q want %q", got, want)
		}
	}

	t.Run("test", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)

	})

	t.Run("unknown", func(t *testing.T) {

		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}