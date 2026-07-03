package main

import "fmt"

func sendData(ch chan<- int) {
	ch <- 10
	close(ch)
}

func receiveData(ch <-chan int) {
	for value := range ch {
		fmt.Println(value)
	}
}

func Implement() {
	ch := make(chan int)

	go sendData(ch)

	receiveData(ch)
}
