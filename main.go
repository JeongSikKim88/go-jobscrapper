package main

import (
	"strings"
	"fmt"
)

func lenAndUpper(name string) (int, string){
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, upperName := lenAndUpper("jack")
	fmt.Println(totalLength,upperName)
}
