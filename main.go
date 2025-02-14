package main

import (
	"fmt"
	"golangweek2github/mypkg"
)

func main() {
	message := mypkg.Welcome()
	fmt.Println(message)
}
