package main

import (
	"fmt"
)

func ClosingChennel() {

	ch := make(chan int, 5)
	go func() {

		ch <- 100
		ch <- 200
		ch <- 300
		ch <- 400
		ch <- 500
		close(ch)

	}()
	value, ok := <-ch
	fmt.Println(value, ok)
	for values := range ch {
		fmt.Println("values from the range is ", values)

	}

}
