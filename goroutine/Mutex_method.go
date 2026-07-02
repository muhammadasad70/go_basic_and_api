package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func IncrementNum(wg *sync.WaitGroup) {
	defer wg.Done()
	defer mu.Lock()
	counter++
	defer mu.Unlock()
}

func Mutex_Method() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(2)

		go Increment(&wg)
		go Increment(&wg)
	}
	wg.Wait()

	fmt.Println("Final counter :", counter)
}
