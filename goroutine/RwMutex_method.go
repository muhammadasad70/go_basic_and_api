package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	rw sync.RWMutex
	wg sync.WaitGroup
)

func Read(id int) {
	defer wg.Done()

	rw.RLock()

	fmt.Printf("Reader %d is reading: %d\n", id, counter)

	time.Sleep(2 * time.Second)

	rw.RUnlock()
}

func Writer() {
	defer wg.Done()

	rw.Lock()
	fmt.Println("Writer is updating the record so wait till")
	counter++

	time.Sleep(2 * time.Second)
	fmt.Println("Counter updated to ", counter)

	rw.Unlock()
}

func ReadWrite_Method() {
	wg.Add(4)

	go Read(1)
	go Read(2)
	go Read(3)

	time.Sleep(500 * time.Millisecond)

	go Writer()

	wg.Wait()
}
