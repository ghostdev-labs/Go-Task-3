package ATM_Machine

import (
	"fmt"
	"log"
	"os"
)

func NewLine(numberOfLines int) {
	for i := 0; i < numberOfLines; {
		fmt.Printf("\n")
		i++
	}
}

func ErrorChecker(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Exit() {
	NewLine(1)
	fmt.Println("\nThanks for using our service!!! \nHave a nice day")
	os.Exit(0)
}

func anotherTransaction() {
	fmt.Printf("\n\nDo you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
	var transaction int
	_, err := fmt.Scan(&transaction)
	ErrorChecker(err)

	switch transaction {
	case 1:
		Menu()
	case 2:
		Exit()
	default:
		fmt.Println("ERROR: invalid input, choose a valid option")
		NewLine(1)
		anotherTransaction()
	}
}

func Menu() {
	NewLine(1)
	fmt.Println("---------- MENU ----------")
	NewLine(1)
	fmt.Println("1. Deposit \t 2. Check Balance")
	fmt.Println("3. Withdraw \t 4. Change PIN")
	fmt.Println("0. Exit")
	NewLine(1)

	fmt.Print("Enter a number: ")
	var menuNumber int
	_, err := fmt.Scan(&menuNumber)
	ErrorChecker(err)

	switch menuNumber {
	case 1:
		Deposit()
	case 2:
		CheckBalance()
	case 3:
		Withdraw()
	case 4:
		ChangePIN()
	case 0:
		Exit()
	default:
		fmt.Println("ERROR: invalid input, select a valid operation")
		NewLine(1)
		Menu()
	}
}

var accountBalance float64

func Deposit() {
	NewLine(1)
	fmt.Println("--------------DEPOSIT----------------")
	NewLine(1)

	var depositAmount float64

	fmt.Println("How much do you want to deposit?")
	NewLine(1)

	fmt.Print("Enter amount: ")

	for {
		_, err := fmt.Scan(&depositAmount)
		ErrorChecker(err)
		if depositAmount > 0 {
			accountBalance += depositAmount
			fmt.Printf("You have successfully deposited ₦%.2f", depositAmount)
			NewLine(1)
			break
		} else {
			fmt.Printf("ERROR: cannot deposit ₦%.2f, enter valid amount: ", depositAmount)
		}
	}
	anotherTransaction()
}

func CheckBalance() {
	NewLine(1)
	fmt.Println("----------BALANCE----------")
	fmt.Println("Your account balance is: ", accountBalance)
	NewLine(1)
	anotherTransaction()
}

func Withdraw() {
	NewLine(1)
	fmt.Println("--------------WITHDRAWAL----------------")
	NewLine(1)

	var withdrawalAmount float64
	var menuNumber int

	fmt.Print(" How much would you like to withdraw?")
	NewLine(1)
	fmt.Println("1. 2000 \t\t 2. 5000")
	fmt.Println("3. 10000 \t\t 4. others")
	NewLine(1)

	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&menuNumber)
	NewLine(1)
	ErrorChecker(err)

	for {
		if withdrawalAmount > accountBalance {
			fmt.Println("Insufficient Funds")
			anotherTransaction()
		}

		switch menuNumber {
		case 1:
			accountBalance -= 2000
			fmt.Println("You have successfully withdrawn ₦2000")
		case 2:
			accountBalance -= 5000
			fmt.Println("You have successfully withdrawn ₦5000")
		case 3:
			accountBalance -= 10000
			fmt.Println("You have successfully withdrawn ₦10000")
		case 4:
			fmt.Print("Enter the amount you wish to withdraw: ")
			for {
				_, err := fmt.Scan(&withdrawalAmount)
				ErrorChecker(err)
				if withdrawalAmount > accountBalance {
					fmt.Println("Insufficient Funds")
					anotherTransaction()
				} else if withdrawalAmount > 0 {
					accountBalance -= withdrawalAmount
					fmt.Printf("You have successfully withdrawn ₦%.2f", withdrawalAmount)
					NewLine(1)
					break
				}
			}
		default:
			fmt.Println("ERROR: invalid input, select a valid option")
		}

		anotherTransaction()

	}
}
