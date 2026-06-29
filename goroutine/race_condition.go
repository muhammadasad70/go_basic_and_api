package main

import (
	"fmt"
	"sync"
)

var counter int = 0

func Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter++
}

func Method() {
	var wg sync.WaitGroup

	wg.Add(2)
	go Increment(&wg)
	go Increment(&wg)

	wg.Wait()
	fmt.Println(counter)

}
