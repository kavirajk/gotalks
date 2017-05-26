package main

import (
	"fmt"

	"github.com/kavirajk/gopl-ex/ch9/bank"
)

//START OMIT
func main() {
	go func() {
		bank.Deposit(100) // HL
		fmt.Println(bank.Balance())
	}()

	go bank.Deposit(300) // HL
}

// END OMIT
