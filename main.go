package main

import (
	"github.com/JeongSikKim88/go-jobscrapper/accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("jack")
	fmt.Println(account)
}
