package main

import (
	"fmt"
	"time"
)

func FetchUserData(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Loading.........."
}

func UserImplement() {
	ch := make(chan string)

	go FetchUserData(ch)
	select {
	case ms := <-ch:
		fmt.Println(ms)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")

	}
}
