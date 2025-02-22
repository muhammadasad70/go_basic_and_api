package main

import (
	"fmt"
)

// func main() {
// 	tasks()
// 	ForStu()
// 	Unpacking()

// }

// this is for the task 7
type Product struct {
	id    int
	title string
	price int
}

func ForStu() {
	// description of this task is that it will make a array of the structure and also adding values in the structure and
	// also use for the displaying the values we can make the array of the structure
	product := []Product{
		Product{10, "Bread", 120},
		Product{20, "Meet", 150},
	}
	fmt.Println("Task 7 Output :")
	fmt.Println(product)
	product = append(product, Product{30, "Lays", 50})
	fmt.Println(product)
}

func tasks() {
	// Task 1: Print hobbies
	hobbies := [3]string{"watching Netflix :", "Cooking :", "Watching Reels :"}
	fmt.Println("Task 1 output:")
	for i := 0; i < len(hobbies); i++ {
		fmt.Println(i, ":", hobbies[i])
	}

	// Task 2: Print specific hobby
	fmt.Println("Task 2 output:")
	fmt.Println(hobbies[0])

	// Task 2 (continued): Slicing hobbies
	fmt.Println("Task 2 output also:")
	slice1 := hobbies[1:]
	fmt.Println(slice1)

	// Task 3: Print another slice
	slice2 := hobbies[0:2]
	fmt.Println("Task 3 output:")
	fmt.Println(slice2)

	// Task 4: Replace the 2nd and 3rd elements
	copy(slice2, hobbies[1:])
	fmt.Println("Task 4 output:")
	fmt.Println(slice2)

	// Task 5: Define a dynamic slice (instead of an array)
	goals := []string{"I want to be a Go developer:", "I also want to be a trader:"}
	fmt.Println("Task 5 output:")
	for j := 0; j < len(goals); j++ {
		fmt.Println(j, ":", goals[j])
	}

	// Task 6: Modify the second goal and append a new one
	goals[1] = "I want to be a good trader"
	fmt.Println("Task 6 output:")
	fmt.Println(goals)

	// Append new goals
	goals = append(goals, "I want to be a Golang developer:")
	fmt.Println("Updated goals:")
	fmt.Println(goals)
	// this is for the unpacking of the list element

}
func Unpacking() {
	values := []int{10, 20, 30, 40, 50}
	fmt.Println(values)
	AddingValues := []int{60, 70, 80, 90, 100}
	Updated := append(values, AddingValues...)
	// AddingValues... (this .... mean that it will include all the values of the list)
	fmt.Println(Updated)

}
