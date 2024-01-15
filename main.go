package main

import "fmt"

var balance float64 = 0

func main() {
	fmt.Println(getBalance())
	deposit(1000)
	fmt.Println(getBalance())
	withdraw(200)
	fmt.Println(getBalance())
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
