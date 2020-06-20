package main

import (
	"fmt"

	"github.com/JeongSikKim88/go-jobscrapper/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "first")
	err := dictionary.Update(baseWord, "second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
}
