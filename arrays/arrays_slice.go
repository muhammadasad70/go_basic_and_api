package main

import (
	"fmt"
)

/*
	Consider the basic syntax of the array

example is:
pricce:= .... it is just like the declearing the variable but actually we also use it for the decelereation of the array
price:=[4]......[4] mean that we will declear the length of the array which is 4 so mean
in this [] it will take the size or length of the araray
price:=[4]float64 ......float64 this will indicate the type of the values we will initialized in this
price:=[4]float{}....{}mean that the values taken by the arrays
*/
func CreatingArray() {
	price := [4]int{23, 45, 56, 78}
	for index, value := range price {
		fmt.Println("Index is :", index)
		fmt.Println("Value is :", value)
	}
	price[2] = 90
	// fmt.Println(price)
	// specialprice := price[1:3]
	// fmt.Println(specialprice)
	// highlitedprice := specialprice[0:]

	// fmt.Println(highlitedprice)
	// this is for the iteration on the slice
	// for i := 0; i <= 3; i++ {
	// 	fmt.Println(price[i])
	// }
	// this is the true method to get all the values by the for loops
}
func makingslice() {
	values := [5]int{10, 20, 30, 40, 50}
	fmt.Println(values)
	slice1 := values[2:5]
	slice2 := slice1[1:]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

}

func main() {
	CreatingArray()
	// makingslice()
}
