package main

import "fmt"

var balance float64 = 0

func main() {
	fmt.Printf("The current balance is %.02f\n", getBalance())
	deposit(1000)
	fmt.Printf("Balance after deposit of % .2f%.02f\n", getBalance())
	withdraw(200)
	fmt.Printf("Balance after the withdrawal %.02f\n", getBalance())
}

func getBalance() float64 {
	return balance
}

func deposit(m float64) {
	balance += m
}

func withdraw(m float64) {
	balance -= m
}
