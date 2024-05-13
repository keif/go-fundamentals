package main

import (
	"example.com/bank/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var choice int
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Printf("Error getting account balance: %v\n", err)
		fmt.Println("----")
		//panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7! ", randomdata.PhoneNumber())

	for {
		presentOptions()

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

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
			fileops.WriteBalanceToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteBalanceToFile(accountBalance, accountBalanceFile)
			fmt.Println("Your balance is now: ", accountBalance)
		default:
			fmt.Println("Goodbye! ")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}
