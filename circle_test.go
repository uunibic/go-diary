package main

import "testing"

func TestArea(t *testing.T) {

	t.Run("Area of Rectangle", func(t *testing.T) {

		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("Area of Circle", func(t *testing.T) {

		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
