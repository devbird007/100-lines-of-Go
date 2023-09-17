package main

import (
	"fmt"
)

func withdraw(bal *float64, price float64) {
	*bal = *bal - price
}

func main() {
	var acctBal float64 = 3000.43
	var price float64 = 1000

	// Create a pointer to account balance
	acctBalPtr := &acctBal
	fmt.Printf("Account balance is %.2f\n", acctBal)

	withdraw(acctBalPtr, price)
	fmt.Printf("Account balance is %.2f\n", acctBal)
}
