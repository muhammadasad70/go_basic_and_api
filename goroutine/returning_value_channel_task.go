package main

import (
	"fmt"
)

func GenerateNumber(ch chan int) {

	number := 50

	ch <- number
}

func Task_Method() {
	ch := make(chan int)

	go GenerateNumber(ch)

	value := <-ch

	fmt.Println("Value returned by the channels is :", value)
}
