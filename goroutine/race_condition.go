package main

import (
	"fmt"
	"sync"
)

var counter int = 0

func Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}

func RaceCondition_Method_1() {
	var wg sync.WaitGroup

	wg.Add(2)
	go Increment(&wg)
	go Increment(&wg)

	wg.Wait()
	fmt.Println(counter)

}
func RaceCondition_Method_2() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(2)

		go Increment(&wg)
		go Increment(&wg)
	}
	wg.Wait()
	fmt.Println("Final Counter", counter)

}
