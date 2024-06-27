package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalance = "balance.txt"
const accountBalanceFile = accountBalance

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000.0, errors.New("failed to read file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000.0, errors.New("failed to parse balance")
	}
	return balance, nil

}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountbalance, error = getBalanceFromFile()
	if error != nil {
		fmt.Println("Failed to get balance from file.")
		fmt.Println("Err")
	}

	fmt.Println("Welcome to the Bank")

	for {

		fmt.Println("what do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your account balance is: ", accountbalance)
		} else if choice == 2 {
			fmt.Print("Enter the amount you want to deposit: ")
			var deposit float64
			fmt.Scan(&deposit)
			accountbalance += deposit
			fmt.Println("Your account balance is: ", accountbalance)
			writeBalanceToFile(accountbalance)

		} else if choice == 3 {
			fmt.Print("Enter the amount you want to withdraw: ")
			var withdraw float64
			fmt.Scan(&withdraw)
			if withdraw > accountbalance {
				fmt.Println("Insufficient funds")
			} else {
				accountbalance -= withdraw
				fmt.Println("Your account balance is: ", accountbalance)
			}
		} else if choice == 4 {
			fmt.Println("Thank you for banking with us")
			return
		} else {
			fmt.Println("Invalid choice")
		}
	}

}
