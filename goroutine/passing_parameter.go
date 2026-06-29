package main

import (
	"fmt"
	"sync"
)

func PrintNumber(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(num)
}
