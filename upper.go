package main

import "strings"
import "fmt"

func Upper(input string) string {
	return strings.ToUpper(input)
}

func main() {
    original := "hello, world"
    uppercased := Upper(original)
    fmt.Println(uppercased)
}
