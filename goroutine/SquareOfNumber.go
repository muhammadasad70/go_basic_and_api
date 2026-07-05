package main

import (
	"fmt"
)

func Square(num_val int, ch chan int) {
	square := num_val * num_val

	ch <- square
}

func SquareImplement() {
	ch := make(chan int)

	fmt.Println("Enter the value which square you want to print or get")
	val := 0
	fmt.Scan(&val)

	go Square(val, ch)
	sqr := <-ch

	fmt.Println("Square of the entered function is :", sqr)
}
