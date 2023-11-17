package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestUpper(t *testing.T) {
	input := "golang"
	expected := "GOLANG"
	result := strings.ToUpper(input)

	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
	}
}

func BenchmarkUpper(b *testing.B) {
	input := "benchmark"
	for i := 0; i < b.N; i++ {
		Upper(input)
	}
}

func ExampleUpper() {

	input := "i am mahesh"
	result := Upper(input)

	fmt.Println(result)

	// Output: I AM MAHESH
}