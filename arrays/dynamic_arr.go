package main

import "fmt"

// dynamic slice mean that we use the different functions on the slice like the append
func dynamicarray() {
	price := []int{20, 30, 40}
	fmt.Println(price)
	updatedprice := append(price, 50)
	fmt.Println(updatedprice)

}

// func main() {
// 	dynamicarray()

// }
