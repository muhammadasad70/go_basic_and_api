package main

import "fmt"

func main() {

	//  Create Map
	students := make(map[string]int)

	//  Insert Data
	students["Asad"] = 90
	students["Ali"] = 85
	students["Ahmed"] = 95

	fmt.Println("After Insertion:")
	fmt.Println(students)

	//  Access Value

	fmt.Println("\nMarks of Asad:")
	fmt.Println(students["Asad"])

}
