package main

import (
	"fmt"
)

// (when we have to use the maps)
// we use it when we have to represent the collection of the releated proporties

// func main() {
// 	Formap()
// 	AddingValuesToMap()

// }

// maps in go is like the list in others language and also like a hashing table in other programmings language
// like in list we have the value with it index so map is similar to the list in which we have the values with keys
// we can directly access with keys
// like a array
// syntax for the map is
// colors:=map[string]........[string]mean that all the keys for the values is string
// colors:=map[string]string{}.....2nd string mean the values in the map is type string
func Formap() {
	colors := map[string]string{
		"red":   "1213",
		"green": "1324",
	}
	// this is for accessing the values for the keys
	fmt.Println(colors["red"])
	for index, value := range colors {
		fmt.Println("Index is :", index)
		fmt.Println("Value is :", value)
	}

}
func AddingValuesToMap() {
	colors := map[string]string{}
	colors["Blue"] = "1357"
	fmt.Println(colors)
	// andother map with int type keys
	// price := make(map[int]string,2)in this we are also mention the capacity of the map
	price := make(map[int]string)
	price[10] = "Bread"
	price[20] = "Lays"
	fmt.Println(price)
	// this is for deleting map with value
	// delete(price, 20)
	// fmt.Println(price)
	fmt.Println(price)
}
