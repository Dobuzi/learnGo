package main

import (
	"fmt"

	"github.com/Dobuzi/learnGo/accounts"
)

func main() {
	jwAccount := accounts.MakeAccount("JW")
	jwAccount.Deposit(100)
	fmt.Println(jwAccount.ShowBalance())
	err := jwAccount.Withdraw(101)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jwAccount.ShowBalance())
}
