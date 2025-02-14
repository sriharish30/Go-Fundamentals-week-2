package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func multiply(a int, b int) int {
	return a * b
}
func main() {
	fmt.Println("sum:", add(10, 5))
	fmt.Println("multiply:", multiply(10, 5))
}
