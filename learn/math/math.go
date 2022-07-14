package math

import "fmt"

func init() {
	fmt.Println("In math init")
}

func Add(a, b int) int {
	return a + b
}
func Subtract(a, b int) int {
	return a - b
}