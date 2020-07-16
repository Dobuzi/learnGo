package main

import (
	"fmt"

	"github.com/Dobuzi/learnGo/accounts"
)

func main() {
	jwAccount := accounts.MakeAccount("JW")
	jwAccount.Deposit(100)
	jwAccount.ChangeOwner("Dobuzi")
	fmt.Println(jwAccount)
}
