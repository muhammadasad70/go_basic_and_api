package main

import (
	"fmt"
)

func Counter() {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			ch <- 1

		}()

	}

	count := 0

	for i := 0; i < 10; i++ {
		count = count + <-ch
	}

	fmt.Println("Final Counter is :", count)
}
