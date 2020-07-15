package accounts

// Account : banking account
type Account struct {
	owner   string
	balance int
}

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
