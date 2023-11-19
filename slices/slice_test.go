package main

import (
	"testing"
	"fmt"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T){
		numbers := []int{1,2,3,4,5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("expected %d but got %d given, %v", want, got, numbers)
		}
	})

	// t.Run("collection of any size", func(t *testing.T) {
	// 	numbers := []int{1,2,3}

	// 	got := Sum(numbers)
	// 	want := 6

	// 	if got != want {
	// 		t.Errorf("expected %d but got %d given, %v", want, got, numbers)
	// 	}
	// })
}

func BenchmarkSum(b *testing.B) {

	numbers := []int{1,2,3,4,5}

	for i := 0; i < b.N; i++ {

		Sum(numbers)
	}
}

func ExampleSum() {

	numbers := []int{1,2,3,4,5}
	sum := Sum(numbers)
	fmt.Println(sum)

	// Output: 15
}

