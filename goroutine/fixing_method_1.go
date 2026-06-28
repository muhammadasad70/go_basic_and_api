//so with the problem of the goroutine we have two ways of solution
// --> 1 time.Sleep(1*time.second)
// --> 2 WaitGroup

// Method 1

package main

import (
	"fmt"
	"time"
)

func Hello_2() {
	fmt.Println("Hello")

}
func WithSleepMethod() {
	go Hello_2()
	time.Sleep(1 * time.Second)
}
