package main

import (
	"fmt"
	"os"
)

// as this is the variable that is storing the value but for storing the value in the file
// we can write the write function to write the data in the file
func WriteTheBalanceToFile(balance int) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

var globalBalance int = 500000

func Main_For_Bank() {
	for {
		if !Process() {
			break

		}

	}

}

func Welcome() {
	fmt.Println("Welcome to go bank:")
	fmt.Println("Select what you want to do")
	fmt.Println("Press 1 for check balance:")
	fmt.Println("Press 2 for deposite money:")
	fmt.Println("Press 3 for withdraw money:")
	fmt.Println("Press 4 for exit")
}
func UserInputs() int {
	var choice int
	fmt.Scan(&choice)
	fmt.Println("Your choice is :", choice)
	return choice
}
func Process() bool {
	Welcome()
	var choice int = UserInputs()
	switch choice {
	case 1:
		AccountBalance()
	case 2:
		Deposite()
	case 3:
		Withdraw()
	default:
		fmt.Println("Good Bye ")
		return false

	}
	return true

}
func AccountBalance() {
	fmt.Println("Your Account Balance is :", globalBalance)
}
func Deposite() {
	var old_balance int = globalBalance
	fmt.Println("Your current balance is :", old_balance)
	var deposite int
	fmt.Println("Enter the amount you want to deposite ")
	fmt.Scan(&deposite)
	var newBalance int = old_balance + deposite
	WriteTheBalanceToFile(newBalance)
	// fmt.Println("Your new balance is :", globalBalance)
}
func Withdraw() {
	var withdraw int
	fmt.Println("Your current balance is:", globalBalance)
	fmt.Println("Enter the amount you want to withdraw:")
	fmt.Scan(&withdraw)
	if withdraw > globalBalance {
		fmt.Println("Insufficient funds!")
	} else {
		globalBalance -= withdraw
		newBalance := globalBalance
		WriteTheBalanceToFile(newBalance)
		// fmt.Println("Your new balance is:", globalBalance)
	}
}

// func main() {
// 	Main_For_Bank()
// }
