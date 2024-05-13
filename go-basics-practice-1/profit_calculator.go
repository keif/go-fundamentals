package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
//  1. Validate user input
//     => show error message anbd exit if invalid input is provided
//     - no negative numbers
//     - not 0
//  2. store calculated results into a file
func main() {
	revenue, err := getUserInput("Enter revenue: ")
	if err != nil {
		fmt.Printf("Error getting revenue: %v\n", err)
		fmt.Println("----")
		return
	}
	expenses, err := getUserInput("Enter expenses: ")
	if err != nil {
		fmt.Printf("Error getting expenses: %v\n", err)
		fmt.Println("----")
		return
	}
	taxRate, err := getUserInput("Enter tax rate: ")
	if err != nil {
		fmt.Printf("Error getting tax rate: %v\n", err)
		fmt.Println("----")
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("(EBT) Earnings Before Tax: %.2f\n", ebt)
	fmt.Printf("Profit:  %.2f\n", profit)
	fmt.Printf("Ratio:  %.2f\n", ratio)
	saveFinancialsToFile(ebt, profit, ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be greater than zero")
	}
	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	// Earnings Before Tax
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func saveFinancialsToFile(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("financials.txt", []byte(results), 0644)
}
