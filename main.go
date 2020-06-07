package main

import (
	"fmt"
)

func main() {
	names := []string{"jack", "gonyon", "surl"}
	names = append(names, "leo")
	fmt.Println(names)
}
