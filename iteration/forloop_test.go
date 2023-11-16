package main

import "testing"

func TestForloop(t *testing.T) {

	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

}