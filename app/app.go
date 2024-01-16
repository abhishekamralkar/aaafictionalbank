package app

var balance float64 = 0

func Balance() float64 {
	return balance
}

func Deposit(m float64) {
	balance += m
}

func Withdraw(m float64) {
	balance -= m
}
