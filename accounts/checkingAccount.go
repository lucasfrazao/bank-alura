package accounts

import "bank-alura/customers"

type CheckingAccount struct {
	Holder                      customers.Holder
	AgencyNumber, AccountNumber int
	balance                     float64
}

func (c *CheckingAccount) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.balance

	if canWithdraw {
		c.balance -= withdrawValue
		return "Withdrawal successful!"
	} else {
		return "Withdrawal not possible!"
	}
}

func (c *CheckingAccount) CashDeposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.balance += depositValue
		return "Deposit successful!", c.balance
	} else {
		return "Deposit not possible!", c.balance
	}
}

func (c *CheckingAccount) Transfer(transferValue float64, accountTo *CheckingAccount) bool {
	if transferValue < c.balance && transferValue > 0 {
		c.balance -= transferValue
		accountTo.CashDeposit(transferValue)
		return true
	} else {
		return false
	}
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}
