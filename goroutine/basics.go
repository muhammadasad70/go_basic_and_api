package main

import (
	"fmt"
)

func Hello() {
	fmt.Println("Hello")
}

// go hello()  go is the keyword use to create the goroutine
// fmt.Println("Main")
//so in this case the output is the main because go hello()take time to execute so main not wait for it to run and finised so it will continue printing
// so main continue and print the main and exit so in this case the hello()never execute and killed
