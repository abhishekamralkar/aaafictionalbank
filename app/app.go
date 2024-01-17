package app

import "fmt"

var balance float64 = 0

func Balance() float64 {
	return balance
}

func Deposit(m float64) float64 {
	balance += m
	return balance
}

func Withdraw(m float64) float64 {
	if balance < m {
		fmt.Println("Insufficient Funds")
	}
	balance -= m
	return balance
}
