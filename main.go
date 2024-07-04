package main

import (
	"bank-alura/accounts"
	"fmt"
)

type verifyAccount interface {
	Withdraw(withdrawValue float64) string
}

func PayBill(account verifyAccount, billValue float64) {
	account.Withdraw(billValue)
}

func main() {
	exampleAccount := accounts.SavingAccount{}
	exampleAccount.CashDeposit(100)
	PayBill(&exampleAccount, 60)

	fmt.Println(exampleAccount.GetBalance())

	luizaAccount := accounts.CheckingAccount{}
	luizaAccount.CashDeposit(500)
	PayBill(&luizaAccount, 100)

	fmt.Println(luizaAccount.GetBalance())
}
