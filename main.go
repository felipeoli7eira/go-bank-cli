package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"errors"
)

const GoBankBalanceFileName = "GoBankBalance.txt"

func execClearOSCommand() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
	}
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(GoBankBalanceFileName, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	file, err := os.ReadFile(GoBankBalanceFileName)

	if err != nil {
		return 0, errors.New("Failed to find balance file.")
	}

	fileContentString := string(file)
	balaceFloat64, err := strconv.ParseFloat(fileContentString, 64)

	if err != nil {
		return 0, errors.New("Feiled to parse stored balance value.")
	}

	return balaceFloat64, nil
}

func main() {
	execClearOSCommand()

	balance, err := getBalanceFromFile()

	welcome()

	if err != nil {
		fmt.Println("❗ Balance history not found")
		fmt.Println()
	}

	for {

		displayOptions()

		var choice int
		fmt.Scan(&choice)

		switch choice {
			case 1:
				execClearOSCommand()
				fmt.Println("🟢 Your balance is:", balance)
				fmt.Println()
				continue
			case 2:
				execClearOSCommand()
				var deposit float64

				fmt.Print("Your deposit: ")
				fmt.Scan(&deposit)

				if deposit <= 0 {
					execClearOSCommand()
					fmt.Println("Invalid deposit amount. Please enter a positive number.")
					continue
				}

				balance += deposit

				writeBalanceToFile(balance)
				execClearOSCommand()
				fmt.Println("Balance updated 🎉! New amount:", balance)
				fmt.Println()
				continue
			case 3:
				execClearOSCommand()
				var withdrawalAmount float64

				fmt.Print("Withdrawal amount: ")
				fmt.Scan(&withdrawalAmount)

				if withdrawalAmount <= 0 {
					execClearOSCommand()
					fmt.Println("🔴 Invalid amount. Must be greater than 0.")
					continue
				}

				if withdrawalAmount > balance {
					execClearOSCommand()
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
						writeBalanceToFile(balance)
						execClearOSCommand()

						fmt.Println("Balance updated 📉! New amount:", balance)
						fmt.Println()

						continue
					} else {
						execClearOSCommand()
						fmt.Println("🟢 No withdrawals made!")
						fmt.Println()

						continue
					}
				}

				balance -= withdrawalAmount

				writeBalanceToFile(balance)
				execClearOSCommand()

				fmt.Println("Balance updated 📉! New amount:", balance)
				fmt.Println()
			default:
				fmt.Println("Goodbye! 👋")
				return

		}
	}
}

func displayOptions() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdrawal")
	fmt.Println("4. Exit")
	fmt.Println()

	fmt.Print("Your choice: ")

}

func welcome() {
	fmt.Println("Welcome to GoBank!")
	fmt.Println()
}
