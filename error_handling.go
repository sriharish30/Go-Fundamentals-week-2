package main

import (
	"errors"
	"fmt"
)

func division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divisible by zero")
	}
	return a / b, nil
}
func main() {
	result, err := division(15, 3)
	if err != nil {
		fmt.Println("error:", err)

	} else {
		fmt.Println("result:", result)
	}
	result, err = division(100, 0)
	if err != nil {
		fmt.Println("error:", err)

	} else {
		fmt.Println("result:", result)
	}
}
