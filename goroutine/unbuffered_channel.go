package main

import "fmt"

func Unbuffered_Method() {
	ch := make(chan int)
	go func() {
		fmt.Println("Sending data")
		ch <- 100
		fmt.Println("Done")
	}()
	fmt.Println("Receiving data ")
	value := <-ch

	fmt.Println(value)
}
