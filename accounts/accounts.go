package accounts

import (
	"errors"
	"fmt"
)

// Account : banking account
type Account struct {
	owner   string
	balance int
}

var errorWithdraw = errors.New("Can't withdraw money")

// MakeAccount : for new account with $0
func MakeAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit : amount on your bank account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// ShowBalance : just show balance
func (a Account) ShowBalance() int {
	return a.balance
}

// Withdraw : withdraw your money from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errorWithdraw
	}
	a.balance -= amount
	return nil
}

// ChangeOwner : change the owner
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// ShowOwner : show owner's name
func (a Account) ShowOwner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.ShowOwner(), "'s account\nHas: ", a.ShowBalance())
}
