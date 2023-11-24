package main

import "testing"

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {

		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Area of Rectangle", func(t *testing.T){
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("Area of Circle", func(t *testing.T){
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}