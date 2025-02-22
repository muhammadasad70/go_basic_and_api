package main

import (
	"fmt"
)

func pntr(){
	var Number int = 10
	var ptr = &Number

	fmt.Println("Hi the address of the ptr is:", &ptr)

	fmt.Println("Hi the address of the ptr is:", *ptr)
}
