package main

import (
	"fmt"
	"github.com/JeongSikKim88/go-jobscrapper/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(definition)
}
