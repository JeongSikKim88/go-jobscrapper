package main

import (
	"fmt"

	"github.com/JeongSikKim88/go-jobscrapper/accounts"
)

func main() {
	account := accounts.NewAccount("jack")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
