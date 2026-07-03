package main

import (
	"fmt"
)

func SelectStatement() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "Data from channel_1"
	}()

	go func() {
		ch2 <- "Data from channel_2"

	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Message from ch 1 ", msg1)
	case msg2 := <-ch2:
		fmt.Println("Message from ch 2", msg2)

	}
}
