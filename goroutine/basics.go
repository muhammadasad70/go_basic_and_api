package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello")
}

func main() {
	go hello()

	fmt.Println("Main")
}
