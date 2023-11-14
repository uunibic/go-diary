package main

import "testing"

func TestHello1(t *testing.T) {

	got:= Hello("Chris")
	want:= "Hello, Chris!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}