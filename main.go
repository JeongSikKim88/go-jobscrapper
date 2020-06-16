package main

import (
	"fmt"

	"github.com/JeongSikKim88/go-jobscrapper/accounts"
)

func main() {
	account := accounts.NewAccount("jack")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
}
