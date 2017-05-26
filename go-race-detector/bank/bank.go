package bank

var balance int

func Deposit(amount int) {
	balance += amount
}

func Balance() int {
	return balance
}
