package accounts

import "bank-alura/customers"

type SavingAccount struct {
	Holder                                 customers.Holder
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingAccount) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.balance

	if canWithdraw {
		c.balance -= withdrawValue
		return "Withdrawal successful!"
	} else {
		return "Withdrawal not possible!"
	}
}

func (c *SavingAccount) CashDeposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.balance += depositValue
		return "Deposit successful!", c.balance
	} else {
		return "Deposit not possible!", c.balance
	}
}

func (c *SavingAccount) GetBalance() float64 {
	return c.balance
}
