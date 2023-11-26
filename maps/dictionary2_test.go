package main

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func TestDict(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

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

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {

	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}

func TestAdd(t *testing.T) {

	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	word := "test"
	definition := "this is just a test"

	assertDefinition(t, dictionary, word, definition)

}