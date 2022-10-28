package ATM_Machine

import "fmt"

var PIN = "1234"

func Auth() {
	for {
		if !VerifyPIN(PIN) {
			continue
		}
		break
	}
}

func VerifyPIN(accountPIN string) bool {
	var inputPIN string
	fmt.Printf("ENTER PIN: ")
	_, err := fmt.Scan(&inputPIN)
	ErrorChecker(err)

	if accountPIN != inputPIN {
		fmt.Println("Incorrect PIN")
		return false
	}

	return true
}

func ChangePIN() {
	NewLine(1)
	fmt.Println("-------------CHANGE PIN-------------")
	for {
		fmt.Printf("Enter new PIN: ")
		var newPIN string
		_, err := fmt.Scan(&newPIN)
		ErrorChecker(err)

		if len(newPIN) != 4 {
			fmt.Println("PIN should be 4 characters long")
			continue
		}

		PIN = newPIN
		fmt.Println("PIN updated successfully")
		anotherTransaction()
		break
	}
}
