package main

import (
	"fmt"
	"sync"
)

func WitGroup_Method() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		Hello()
	}()
	wg.Wait()
	fmt.Println("Responding from WaitGroup")
}
