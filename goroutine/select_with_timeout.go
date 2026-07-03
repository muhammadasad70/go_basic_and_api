package main

import (
	"fmt"
	"time"
)

func SelectStatementTimeout() {
	ch1 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Data from channel_1"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Message from ch 1 ", msg1)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")

	}
}
