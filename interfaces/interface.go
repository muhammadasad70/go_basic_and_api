package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}
type englishbot struct{}
type spnishbot struct{}

//	func main() {
//		en := englishbot{}
//		sp := spnishbot{}
//		printGreeting(en)
//		printGreeting(sp)
//	}
func (englishbot) getGreeting() string {
	fmt.Println("hello")
	return "Hi There!"
}
func (spnishbot) getGreeting() string {
	return "Hola!"
}

func printGreeting(eb bot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sp spnishbot) {
// 	fmt.Println(sp.greeting())
// }
