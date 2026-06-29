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

func WitGroup_Method_2() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		Hello_1()
	}()

	go func() {
		defer wg.Done()
		Hello_2()
	}()

	go func() {
		defer wg.Done()
		Hello_3()
	}()
	wg.Wait()
	fmt.Println("printing 3 out of 3")
}
