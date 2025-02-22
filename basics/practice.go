package main

// fmt is the defalut pakage
// main is also a package
// main is the entry point which tell the application that our program is starting form this point
// one module is consist of multiple packages
// for importing the module command is {go mog init example.com/GOLANGUAGE }
// important point is that if you want to accesst function or any things outside the file make it in UperCase
import (
	"fmt"
	"math"
)

func Practice() {
	fmt.Println("Hello and welcome to the course:")
	var investment float64 = 1000
	var expectedReturnrate = 5.5
	var year float64 = 10
	var futureValue = investment * math.Pow(1+expectedReturnrate/100, year)

	fmt.Println("The future value is :", futureValue)

	// we can declear the multiple variable in the same line by following this syntax
	// var investment, year, expectedReturnrate float64 = 1000, 10, 5.5
	// var futureValue = investment * math.Pow(1+expectedReturnrate/100, year)

	// fmt.Println("The future value is :", futureValue)
	// we can also take the value from the user by using this syntax
	// var number float64
	// fmt.Println("Enter youe desire number:")
	// fmt.Scan(&number)
	// fmt.Println("The number enter by the user is :", number)
	// simple calculator for calculating the amount

	// for declear the variable we use the var word
}
