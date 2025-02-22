package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {

	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored balance file.")
	}

	return balance, nil
}

func main() {

	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("------------------")
		panic("Can't confinue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		fmt.Println("You have selected:", choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Enter the amount to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your balance is now:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Enter the amount to withdraw: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Withdrawal amount must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Your balance is now:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}
