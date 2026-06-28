//so with the problem of the goroutine we have two ways of solution
// --> 1 time.Sleep(1*time.second)
// --> 2 WaitGroup

// Method 1

package main

import (
	"time"
)

func WithSleepMethod() {
	go Hello()
	time.Sleep(1 * time.Second)
}
