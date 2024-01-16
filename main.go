package main

import (
	"fmt"

	"github.com/abhishekamralkar/aaafictionalbank/app"
)

func main() {
	fmt.Printf("The current balance is %.02f\n", app.Balance())
	app.Deposit(1000)
	fmt.Printf("Balance after deposit of % .2f%.02f\n", app.Balance())
	app.Withdraw(200)
	fmt.Printf("Balance after the withdrawal %.02f\n", app.Balance())
}
