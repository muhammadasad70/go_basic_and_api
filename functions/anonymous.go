package main

// import (
// 	"fmt"
// )

type transformfun func(int) int

// _ mean that we want to declear but we dont want to use it any where
// as if we want to use the function or call the functions then we use the notation like
// transform ().....but if we want to use the functions as the value then we only use the simple name  of the functions like transform
// func main() {
// 	numbers := []int{1, 2, 3, 4}
// 	double := AnontransformNumber(&numbers, func(num int) int {
// 		return num * 2
// 	})
// 	// triple := transformNumber(&numbers, triple)
// 	fmt.Println("the double number for ", numbers, "is :")
// 	fmt.Println(double)
// 	// fmt.Println("the triple number for ", numbers, "is :")
// fmt.Println()

// }

// as in this we have to use the transform function which accept and return the value but we can decleaer the custom type
// also in this we are using our custom type which is the transforms which we have decleared above as the our type
func AnontransformNumber(numbers *[]int, transform transformfun) []int { //in this the functions is decleared with calling it self
	dnumbers := []int{}
	for _, value := range *numbers {
		dnumbers = append(dnumbers, transform(value))
	}
	return dnumbers
}
