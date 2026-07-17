package main

import (
	"fmt"
)

type BankAccount struct {
	balance float32
}

func NewAccount(bal float32) *BankAccount {
	if bal > 0 {
		return &BankAccount{bal}

	} else {
		fmt.Println("баланс меньше нуля")
		return nil
	}
}
func (b *BankAccount) Deposit(sum float32) {
	b.balance = b.balance + sum
}
func (b BankAccount) Balance() float32 {
	return b.balance
}
func main() {
	bala := NewAccount(245234.5245)
	fmt.Println(bala.Balance())
	bala.Deposit(23423.4234)
	fmt.Println(bala.Balance())
}
