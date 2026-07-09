package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter_1 atomic.Int64

func AtomicOperation(wg *sync.WaitGroup) {
	defer wg.Done()
	counter_1.Add(1)

}

func AtomicImplement() {
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(2)
		go AtomicOperation(&wg)
		go AtomicOperation(&wg)

	}
	wg.Wait()

	fmt.Println("Final Counter is :", counter_1.Load())

}
