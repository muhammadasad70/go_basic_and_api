package main

import (
	"fmt"
)

func For_calculation() {
	revenue, expenses, taxRate := TakeValues()
	EBT, EAT, ratio := Calculation(revenue, expenses, taxRate)
	Display(EBT, EAT, ratio)
}

func TakeValues() (int, int, float64) {
	var revenue, expenses int
	var taxRate float64
	fmt.Println("Enter the revenue please:")
	fmt.Scan(&revenue)
	fmt.Println("Enter the expenses please:")
	fmt.Scan(&expenses)
	fmt.Println("Enter the taxRate please:")
	fmt.Scan(&taxRate)
	return revenue, expenses, taxRate

}
func Calculation(revenue, expenses int, tax_rate float64) (int, float64, float64) {
	var EBT int = revenue - expenses
	var EAT float64 = float64(EBT) * (1 - tax_rate)
	var ratio float64 = float64(EBT) / EAT
	return EBT, EAT, ratio
}

func Display(EBT int, EAT float64, ratio float64) {
	fmt.Printf("The earning before tax is %.1d\nThe earning after tax is %.1f\nAnd the ratio is %.1f\n", EBT, EAT, ratio)
}
