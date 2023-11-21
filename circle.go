package main

import "math"

// This is a function:
//-------------------------------------------
// func Area(rectangle Rectangle) float64{
// 	return rectangle.Width * rectangle.Height
// }
//-------------------------------------------

// This is a struct:

type Rectangle struct {
	Width float64
	Height float64
}

// This is a method:

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}