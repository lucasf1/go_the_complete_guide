package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println("Error: ", err)
		// panic("Can't confinue, sorry.")
		return
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println("Error: ", err)
		// panic("Can't confinue, sorry.")
		return
	}

	tax_rate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println("Error: ", err)
		// panic("Can't confinue, sorry.")
		return
	}

	// ebt := revenue - expenses
	// profit := ebt * (1 - tax_rate/100)
	// ratio := ebt / profit
	ebt, profit, ratio := calculate(revenue, expenses, tax_rate)

	fmt.Printf("EBT: %0.1f\n", ebt)
	fmt.Printf("Profit: %0.1f\n", profit)
	fmt.Printf("Ratio: %0.3f\n", ratio)

	storeResults(ebt, profit, ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Input must be greater than 0")
	}

	return userInput, nil
}

func calculate(revenue, expenses, tax_rate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - tax_rate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}

func storeResults(ebt, profit, ratio float64) {
	resuts := fmt.Sprintf("EBT: %0.1f\nProfit: %0.1f\nRatio: %0.3f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(resuts), 0644)
}
