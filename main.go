package main

import (
	"fmt"
	"github.com/ghostdev-labs/go-task-3/ATM-Machine"
)

func Welcome() {
	fmt.Println("---------- WELCOME TO ALTSCHOOL ATM ----------")
	ATM_Machine.NewLine(1)
}

func main() {
	Welcome()
	ATM_Machine.Auth()
	ATM_Machine.Menu()
}
