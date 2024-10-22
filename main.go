package main

import (
	"fmt"
	"github.com/felipeoli7eira/fileops"
	"github.com/felipeoli7eira/terminal"
	"github.com/Pallinder/go-randomdata"
)

const GoBankBalanceFileName = "GoBankBalance.txt"

func main() {
	terminal.ExecClearOSCommand()

	balance, err := fileops.GetFloatFromFile(GoBankBalanceFileName)

	welcome()

	if err != nil {
		fmt.Println("❗ Balance history not found")
		fmt.Println()
	}

	for {

		fmt.Println("What do you want to do?")
		presentationOptions()

		fmt.Print("Your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			terminal.ExecClearOSCommand()
			fmt.Println("🟢 Your balance is:", balance)
			fmt.Println()
			continue
		case 2:
			terminal.ExecClearOSCommand()
			var deposit float64

			fmt.Print("Your deposit: ")
			fmt.Scan(&deposit)

			if deposit <= 0 {
				terminal.ExecClearOSCommand()
				fmt.Println("Invalid deposit amount. Please enter a positive number.")
				continue
			}

			balance += deposit

			fileops.WriteFloatToFile(balance, GoBankBalanceFileName)
			terminal.ExecClearOSCommand()
			fmt.Println("Balance updated 🎉! New amount:", balance)
			fmt.Println()
			continue
		case 3:
			terminal.ExecClearOSCommand()
			var withdrawalAmount float64

			fmt.Print("Withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				terminal.ExecClearOSCommand()
				fmt.Println("🔴 Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawalAmount > balance {
				terminal.ExecClearOSCommand()
				fmt.Println("🔴 Your withdrawal amount is greater than your available balance. Do you want to withdraw your entire available balance?")

				fmt.Println()
				fmt.Println("Y. Yes")
				fmt.Println("N. No")

				var decision string
				fmt.Println()
				fmt.Print("Your choice: ")
				fmt.Scan(&decision)

				if decision == "Y" {
					balance = 0
					fileops.WriteFloatToFile(balance, GoBankBalanceFileName)
					terminal.ExecClearOSCommand()

					fmt.Println("Balance updated 📉! New amount:", balance)
					fmt.Println()

					continue
				} else {
					terminal.ExecClearOSCommand()
					fmt.Println("🟢 No withdrawals made!")
					fmt.Println()

					continue
				}
			}

			balance -= withdrawalAmount

			fileops.WriteFloatToFile(balance, GoBankBalanceFileName)
			terminal.ExecClearOSCommand()

			fmt.Println("Balance updated 📉! New amount:", balance)
			fmt.Println()
		default:
			fmt.Println("Goodbye! 👋")
			return

		}
	}
}

func welcome() {
	fmt.Println("Welcome to GoBank!")
	fmt.Println("Reach us 24/7:", randomdata.PhoneNumber())
	fmt.Println()
}
