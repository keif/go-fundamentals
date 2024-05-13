package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func main() {
	var choice int
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Printf("Error getting account balance: %v\n", err)
		fmt.Println("----")
		return
		//panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		//checkBalance := choice == 1
		//depositMoney := choice == 2
		//withdrawMoney := choice == 3
		//exitApp := choice == 4

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Println("Your current balance is: ", accountBalance)
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Your deposit must be more than $0")
				continue
			}

			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Your balance is now: ", accountBalance)
		case 3:
			fmt.Println("Your current balance is: ", accountBalance)
			fmt.Print("Your withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Your withdrawal must be more than $0")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than your balance.")
			}

			accountBalance -= withdrawalAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Your balance is now: ", accountBalance)
		default:
			fmt.Println("Goodbye! ")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error:", err)
		return 1000, errors.New("Failed to find balance file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return 1000, errors.New("Failed to parse balance file.")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	os.WriteFile(accountBalanceFile, []byte(fmt.Sprintf("%.2f", balance)), 0644)
}
