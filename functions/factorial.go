package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the number for the factorial:")
	var num int
	fmt.Scan(&num)
	fac := factorial(num)
	fmt.Println("The factorial of the ", num, " is :", fac)
	// this is for the recursion factorial
	refac := recurfactorial(num)
	fmt.Println("The factorial of the ", num, " with recursion is :", refac)
	vari := variadic_sum(10, 20, 30)
	fmt.Println("The sum of the varidic is :", vari)

}

// these is an other functions which is variadic functions
/* in this the concept is that some time we want to use the function and allow the user to enter the desire value
so our functions take all the values and then add up them
but in go we have to just add the .. for the time for which we want to add the values
*/
// variadic functions
func variadic_sum(number ...int) int {
	sum := 0
	for _, values := range number {
		sum = sum + values
	}
	return sum
}

// this is the factorial with the recursion
func recurfactorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * recurfactorial(number-1)
}
func factorial(number int) int {
	result := 1
	for i := 1; i <= number; i++ {
		result = result * i
	}
	return result

}
